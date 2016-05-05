package db

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type Lock int

type Opts struct {
	Start, Stop Idx
	Lock        Lock
	Reverse     bool
	Limit       int
}

type Idx interface {
	Cols() []string
	Vals() []interface{}
}

// UniqueIdx is a tagging interface that indentifies indexes that uniquely identify a row.
// Columns that are UNIQUE or PRIMARY fit this interface.
type UniqueIdx interface {
	Idx
	Unique()
}

type Querier interface {
	Query(query string, args ...interface{}) (Rows, error)
}

type Executor interface {
	Exec(string, ...interface{}) (Result, error)
}

// Result is a clone of database/sql.Result
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows is a clone of database/sql.Rows
type Rows interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(dest ...interface{}) error
}

func Scan(q Querier, name string, opts Opts, cb func(data []byte) error, keyCols []string) error {
	query, queryArgs := buildScan(name, opts, keyCols)
	rows, err := q.Query(query, queryArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp []byte
		if err := rows.Scan(&tmp); err != nil {
			return err
		}
		if err := cb(tmp); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return rows.Close()
}

func buildScan(name string, opts Opts, keyCols []string) (string, []interface{}) {
	var buf bytes.Buffer
	var args []interface{}
	fmt.Fprintf(&buf, `SELECT "data" FROM "%s" `, name)
	if len(opts.Start.Vals()) != 0 || len(opts.Stop.Vals()) != 0 {
		buf.WriteString("WHERE ")
	}
	// WHERE Clause
	if len(opts.Start.Vals()) != 0 {
		cols := opts.Start.Cols()
		vals := opts.Start.Vals()
		if len(vals) > len(cols) {
			panic("More vals than cols")
		}
		// Disjunctive normal form, you nerd!
		// Start always has the last argument be a ">="
		// 1, 2, 3 arg scans look like:
		// ((A >= ?))
		// ((A > ?) OR (A = ? AND B >= ?))
		// ((A > ?) OR (A = ? AND B > ?) OR (A = ? AND B = ? AND C >= ?))
		var ors []string
		for i := 0; i < len(vals); i++ {
			var ands []string
			for k := 0; k < i; k++ {
				var cmp string
				if i == len(vals)-1 && k == len(vals)-1 {
					cmp = ">="
				} else if k == len(vals)-1 {
					cmp = ">"
				} else {
					cmp = "="
				}
				ands = append(ands, fmt.Sprintf(`"%s" %s ?`, cols[k], cmp))
			}
			ors = append(ors, strings.Join(ands, " AND "))
		}
		buf.WriteRune('(')
		buf.WriteString(strings.Join(ors, " OR "))
		buf.WriteString(") ")
	}
	if len(opts.Start.Vals()) != 0 && len(opts.Stop.Vals()) != 0 {
		buf.WriteString("AND ")
	}
	if len(opts.Stop.Vals()) != 0 {
		cols := opts.Stop.Cols()
		vals := opts.Stop.Vals()
		if len(vals) > len(cols) {
			panic("More vals than cols")
		}
		// Stop always has the last argument be a "<"
		// 1, 2, 3 arg scans look like:
		// ((A < ?))
		// ((A < ?) OR (A = ? AND B < ?))
		// ((A < ?) OR (A = ? AND B < ?) OR (A = ? AND B = ? AND C < ?))
		var ors []string
		for i := 0; i < len(vals); i++ {
			var ands []string
			for k := 0; k < i; k++ {
				var cmp string
				if k == len(vals)-1 {
					cmp = "<"
				} else {
					cmp = "="
				}
				ands = append(ands, fmt.Sprintf(`"%s" %s ?`, cols[k], cmp))
			}
			ors = append(ors, strings.Join(ands, " AND "))
		}
		buf.WriteRune('(')
		buf.WriteString(strings.Join(ors, " OR "))
		buf.WriteString(") ")
	}
	// ORDER BY
	// Always order by the primary Key,

	return buf.String(), args
}

var (
	ErrColsValsMismatch = errors.New("db: number of columns and values don't match.")
	ErrNoCols           = errors.New("db: no columns provided")
)

func Insert(exec Executor, name string, cols []string, vals []interface{}) error {
	if len(cols) != len(vals) {
		return ErrColsValsMismatch
	}
	if len(cols) == 0 {
		return ErrNoCols
	}

	valFmt := strings.Repeat("?, ", len(vals)-1) + "?"
	colFmtParts := make([]string, 0, len(cols))
	for _, col := range cols {
		colFmtParts = append(colFmtParts, quoteIdentifier(col))
	}
	colFmt := strings.Join(colFmtParts, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", quoteIdentifier(name), colFmt, valFmt)
	_, err := exec.Exec(query, vals...)
	return err
}

func Delete(exec Executor, name string, key UniqueIdx) error {
	cols := key.Cols()
	vals := key.Vals()
	if len(cols) != len(vals) {
		return ErrColsValsMismatch
	}
	if len(cols) == 0 {
		return ErrNoCols
	}

	colFmtParts := make([]string, 0, len(cols))
	for _, col := range cols {
		colFmtParts = append(colFmtParts, quoteIdentifier(col)+" = ?")
	}
	colFmt := strings.Join(colFmtParts, " AND ")
	query := fmt.Sprintf("DELETE FROM %s WHERE %s LIMIT 1;", quoteIdentifier(name), colFmt)
	_, err := exec.Exec(query, vals...)
	return err
}

func Update(exec Executor, name string, cols []string, vals []interface{}, key UniqueIdx) error {
	if len(cols) != len(vals) {
		return ErrColsValsMismatch
	}
	if len(cols) == 0 {
		return ErrNoCols
	}

	idxCols := key.Cols()
	idxVals := key.Vals()
	if len(idxCols) != len(idxVals) {
		return ErrColsValsMismatch
	}
	if len(idxCols) == 0 {
		return ErrNoCols
	}

	colFmtParts := make([]string, 0, len(cols))
	for _, col := range cols {
		colFmtParts = append(colFmtParts, quoteIdentifier(col)+" = ?")
	}
	colFmt := strings.Join(colFmtParts, ", ")

	idxColFmtParts := make([]string, 0, len(idxCols))
	for _, idxCol := range idxCols {
		idxColFmtParts = append(idxColFmtParts, quoteIdentifier(idxCol)+" = ?")
	}
	idxColFmt := strings.Join(idxColFmtParts, " AND ")

	allVals := make([]interface{}, 0, len(vals)+len(idxVals))
	allVals = append(allVals, vals...)
	allVals = append(allVals, idxVals...)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s LIMIT 1;", quoteIdentifier(name), colFmt, idxColFmt)
	_, err := exec.Exec(query, allVals...)
	return err
}

// quoteIdentifier quotes the PostgreSQL way.  Panics on invalid identifiers.
func quoteIdentifier(ident string) string {
	if strings.ContainsAny(ident, "\"\x00") {
		panic(fmt.Sprintf("Invalid identifier %#v", ident))
	}
	return `"` + ident + `"`
}
