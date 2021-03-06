package db

import (
	"bytes"
	"errors"
	"flag"
	"os"
	"testing"
)

type testIdx struct {
	cols []string
	vals []interface{}
}

func (idx *testIdx) Cols() []string {
	return idx.cols
}

func (idx *testIdx) Vals() []interface{} {
	return idx.vals
}

type testUniqueIdx struct {
	cols []string
	vals []interface{}
}

func (idx *testUniqueIdx) Cols() []string {
	return idx.cols
}

func (idx *testUniqueIdx) Vals() []interface{} {
	return idx.vals
}

func (idx *testUniqueIdx) Unique() {}

type execCap struct {
	query  string
	args   []interface{}
	rows   Rows
	result Result
	err    error
}

func (exec *execCap) Exec(query string, args ...interface{}) (Result, error) {
	exec.query = query
	exec.args = args
	return exec.result, exec.err
}

func (exec *execCap) Query(query string, args ...interface{}) (Rows, error) {
	exec.query = query
	exec.args = args
	return exec.rows, exec.err
}

type testRows struct {
	err, closeErr, scanErr error
	vals                   [][]byte
}

func (rs *testRows) Close() error {
	return rs.closeErr
}

func (rs *testRows) Columns() ([]string, error) {
	panic(nil)
}

func (rs *testRows) Err() error {
	return rs.err
}

func (rs *testRows) Next() bool {
	return len(rs.vals) > 0
}

func (rs *testRows) Scan(dest ...interface{}) error {
	if len(dest) != 1 {
		panic("wrong dest count")
	}
	*(dest[0].(*[]byte)) = rs.vals[0]
	rs.vals = rs.vals[1:]
	return rs.scanErr
}

func TestScanQueryFails(t *testing.T) {
	expected := errors.New("expected")
	exec := &execCap{
		err: expected,
	}
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		return nil
	}, testAdap)

	if err != expected {
		t.Log("Expected error", err)
		t.Fail()
	}
}

var testAdap DBAdapter = &postgresqlAdapter{}

func TestScanQueryEmpty(t *testing.T) {
	exec := &execCap{
		rows: &testRows{},
	}
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		panic("don't call me")
	}, testAdap)

	if err != nil {
		t.Log("Unexpected error", err)
		t.Fail()
	}
}

func TestScanQueryCloseFails(t *testing.T) {
	expected := errors.New("expected")
	exec := &execCap{
		rows: &testRows{
			closeErr: expected,
		},
	}
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		panic("don't call me")
	}, testAdap)

	if err != expected {
		t.Log("Expected error", err)
		t.Fail()
	}
}

func TestScanQueryOneRow(t *testing.T) {
	exec := &execCap{
		rows: &testRows{
			vals: [][]byte{[]byte("bar")},
		},
	}
	var dataCap [][]byte
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		dataCap = append(dataCap, data)
		return nil
	}, testAdap)

	if err != nil {
		t.Log("Unexpected error", err)
		t.Fail()
	}
	if len(dataCap) != 1 || !bytes.Equal(dataCap[0], []byte("bar")) {
		t.Log("Wrong rows", dataCap)
		t.Fail()
	}

	if exec.query != `SELECT "data" FROM "foo";` {
		t.Log("Wrong query", exec.query)
		t.Fail()
	}
	if len(exec.args) != 0 {
		t.Log("Wrong args", exec.args)
		t.Fail()
	}
}

func TestScanQueryMultiRow(t *testing.T) {
	exec := &execCap{
		rows: &testRows{
			vals: [][]byte{[]byte("bar"), []byte("baz")},
		},
	}
	var dataCap [][]byte
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		dataCap = append(dataCap, data)
		return nil
	}, testAdap)

	if err != nil {
		t.Log("Unexpected error", err)
		t.Fail()
	}
	if len(dataCap) != 2 || !bytes.Equal(dataCap[0], []byte("bar")) ||
		!bytes.Equal(dataCap[1], []byte("baz")) {
		t.Log("Wrong rows", dataCap)
		t.Fail()
	}
}

func TestScanScanFails(t *testing.T) {
	expected := errors.New("expected")
	exec := &execCap{
		rows: &testRows{
			vals:    [][]byte{[]byte("bar")},
			scanErr: expected,
		},
	}
	var dataCap [][]byte
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		dataCap = append(dataCap, data)
		return nil
	}, testAdap)

	if err != expected {
		t.Log("Expected error", err)
		t.Fail()
	}
	if len(dataCap) != 0 {
		t.Log("Wrong rows", dataCap)
		t.Fail()
	}
}

func TestScanCallbackFails(t *testing.T) {
	expected := errors.New("expected")
	exec := &execCap{
		rows: &testRows{
			vals: [][]byte{[]byte("bar")},
		},
	}
	var dataCap [][]byte
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		dataCap = append(dataCap, data)
		return expected
	}, testAdap)

	if err != expected {
		t.Log("Expected error", err)
		t.Fail()
	}
	if len(dataCap) != 1 || !bytes.Equal(dataCap[0], []byte("bar")) {
		t.Log("Wrong rows", dataCap)
		t.Fail()
	}
}

func TestScanStopEarly(t *testing.T) {
	expected := errors.New("expected")
	exec := &execCap{
		rows: &testRows{
			vals: [][]byte{[]byte("bar")},
			err:  expected,
		},
	}
	var dataCap [][]byte
	err := Scan(exec, "foo", Opts{Lock: LockNone}, func(data []byte) error {
		dataCap = append(dataCap, data)
		return nil
	}, testAdap)

	if err != expected {
		t.Log("Expected error", err)
		t.Fail()
	}
	if len(dataCap) != 1 || !bytes.Equal(dataCap[0], []byte("bar")) {
		t.Log("Wrong rows", dataCap)
		t.Fail()
	}
}

func TestBuildScan(t *testing.T) {
	s := scanStmt{
		name: "foo",
		buf:  new(bytes.Buffer),
		adap: testAdap,
	}

	query, args := s.buildScan()
	if query != `SELECT "data" FROM "foo" FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 0 {
		t.Log("Should have no args", args)
		t.Fail()
	}
}

func TestBuildScanStart(t *testing.T) {
	s := scanStmt{
		name: "foo",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Start: &testIdx{
				cols: []string{"bar"},
				vals: []interface{}{1},
			},
		},
		adap: testAdap,
	}
	query, args := s.buildScan()
	if query != `SELECT "data" FROM "foo" WHERE (("bar" >= ?)) ORDER BY "bar" ASC FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 1 || args[0] != 1 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestBuildScanStop(t *testing.T) {
	s := scanStmt{
		name: "foo",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Stop: &testIdx{
				cols: []string{"bar"},
				vals: []interface{}{1},
			},
		},
		adap: testAdap,
	}

	query, args := s.buildScan()
	if query != `SELECT "data" FROM "foo" WHERE (("bar" < ?)) ORDER BY "bar" ASC FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 1 || args[0] != 1 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestBuildScanStartStop(t *testing.T) {
	s := scanStmt{
		name: "foo",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Start: &testIdx{
				cols: []string{"bar"},
				vals: []interface{}{1},
			},
			Stop: &testIdx{
				cols: []string{"baz"},
				vals: []interface{}{2},
			},
		},
		adap: testAdap,
	}
	query, args := s.buildScan()
	if query != `SELECT "data" FROM "foo" WHERE (("bar" >= ?)) AND (("baz" < ?))`+
		` ORDER BY "bar" ASC FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 2 || args[0] != 1 || args[1] != 2 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestBuildScanPrefix(t *testing.T) {
	s := scanStmt{
		name: "tab",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Prefix: &testIdx{
				cols: []string{"foo", "bar", "baz", "qux"},
				vals: []interface{}{true, 2},
			},
		},
		adap: testAdap,
	}
	query, args := s.buildScan()
	if query != `SELECT "data" FROM "tab" WHERE "foo" = ? AND "bar" = ?`+
		` ORDER BY "baz" ASC, "qux" ASC FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 2 || args[0] != true || args[1] != 2 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestBuildScanPrefixNoVals(t *testing.T) {
	s := scanStmt{
		name: "tab",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Prefix: &testIdx{
				cols: []string{"foo", "bar"},
				vals: []interface{}{},
			},
		},
		adap: testAdap,
	}
	query, args := s.buildScan()
	if query != `SELECT "data" FROM "tab" ORDER BY "foo" ASC, "bar" ASC FOR SHARE;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 0 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestBuildScanLimitReverseLock(t *testing.T) {
	s := scanStmt{
		name: "foo",
		buf:  new(bytes.Buffer),
		opts: Opts{
			Start: &testIdx{
				cols: []string{"bar"},
				vals: []interface{}{1},
			},
			Limit:   1,
			Reverse: true,
			Lock:    LockNone,
		},
		adap: testAdap,
	}
	query, args := s.buildScan()
	if query != `SELECT "data" FROM "foo" WHERE (("bar" >= ?)) ORDER BY "bar" DESC LIMIT 1;` {
		t.Log("Bad Query", query)
		t.Fail()
	}
	if len(args) != 1 || args[0] != 1 {
		t.Log("Wrong args", args)
		t.Fail()
	}
}

func TestAppendPrefix(t *testing.T) {
	s := scanStmt{
		buf: new(bytes.Buffer),
		opts: Opts{
			Prefix: &testIdx{
				cols: []string{"foo", "bar"},
				vals: []interface{}{1},
			},
		},
		adap: testAdap,
	}

	s.appendPrefix()
	stmt := s.buf.String()
	if stmt != ` WHERE "foo" = ? ORDER BY "bar" ASC` {
		t.Log("Bad Stmt", stmt)
		t.Fail()
	}
	if len(s.args) != 1 || s.args[0] != 1 {
		t.Log("Args didn't match", s.args)
		t.Fail()
	}
}

func TestAppendPrefixAll(t *testing.T) {
	s := scanStmt{
		buf: new(bytes.Buffer),
		opts: Opts{
			Prefix: &testIdx{
				cols: []string{"foo", "bar"},
				vals: []interface{}{1, 2},
			},
		},
		adap: testAdap,
	}

	s.appendPrefix()
	stmt := s.buf.String()
	if stmt != ` WHERE "foo" = ? AND "bar" = ?` {
		t.Log("Bad Stmt", stmt)
		t.Fail()
	}
	if len(s.args) != 2 || s.args[0] != 1 || s.args[1] != 2 {
		t.Log("Args didn't match", s.args)
		t.Fail()
	}
}

func TestAppendOrder(t *testing.T) {
	s := scanStmt{
		buf:  new(bytes.Buffer),
		adap: testAdap,
	}
	s.appendOrder([]string{"bar", "baz"})
	stmt := s.buf.String()
	if stmt != ` ORDER BY "bar" ASC, "baz" ASC` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
}

func TestBuildOrderStmtReverse(t *testing.T) {
	s := scanStmt{
		buf: new(bytes.Buffer),
		opts: Opts{
			Reverse: true,
		},
		adap: testAdap,
	}
	s.appendOrder([]string{"foo"})

	stmt := s.buf.String()
	if stmt != ` ORDER BY "foo" DESC` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
}

func TestBuildStopOneVal(t *testing.T) {
	stmt, args := buildStop([]string{"A", "B"}, []interface{}{1}, testAdap)
	if stmt != `(("A" < ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 1 || args[0] != 1 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestBuildStopTwoVals(t *testing.T) {
	stmt, args := buildStop([]string{"A", "B"}, []interface{}{1, 2}, testAdap)
	if stmt != `(("A" < ?) OR ("A" = ? AND "B" < ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 3 || args[0] != 1 || args[1] != 1 || args[2] != 2 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestBuildStopThreeVals(t *testing.T) {
	stmt, args := buildStop([]string{"A", "B", "C"}, []interface{}{1, 2, 3}, testAdap)
	if stmt != `(("A" < ?) OR ("A" = ? AND "B" < ?) OR ("A" = ? AND "B" = ? AND "C" < ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 6 || args[0] != 1 || args[1] != 1 || args[2] != 2 ||
		args[3] != 1 || args[4] != 2 || args[5] != 3 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestBuildStartOneVal(t *testing.T) {
	stmt, args := buildStart([]string{"A", "B"}, []interface{}{1}, testAdap)
	if stmt != `(("A" >= ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 1 || args[0] != 1 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestBuildStartTwoVals(t *testing.T) {
	stmt, args := buildStart([]string{"A", "B"}, []interface{}{1, 2}, testAdap)
	if stmt != `(("A" > ?) OR ("A" = ? AND "B" >= ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 3 || args[0] != 1 || args[1] != 1 || args[2] != 2 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestBuildStartThreeVals(t *testing.T) {
	stmt, args := buildStart([]string{"A", "B", "C"}, []interface{}{1, 2, 3}, testAdap)
	if stmt != `(("A" > ?) OR ("A" = ? AND "B" > ?) OR ("A" = ? AND "B" = ? AND "C" >= ?))` {
		t.Log("Statement didn't match", stmt)
		t.Fail()
	}
	if len(args) != 6 || args[0] != 1 || args[1] != 1 || args[2] != 2 ||
		args[3] != 1 || args[4] != 2 || args[5] != 3 {
		t.Log("Args didn't match", args)
		t.Fail()
	}
}

func TestInsertWrongColsCount(t *testing.T) {
	err := Insert(nil, "Foo", []string{"one"}, []interface{}{1, 2}, testAdap)

	if err != ErrColsValsMismatch {
		t.Fatal("Expected error, but was", err)
	}
}

func TestInsertNoCols(t *testing.T) {
	err := Insert(nil, "Foo", []string{}, []interface{}{}, testAdap)

	if err != ErrNoCols {
		t.Fatal("Expected error, but was", err)
	}
}

func TestInsertOneVal(t *testing.T) {
	exec := &execCap{}

	Insert(exec, "Foo", []string{"bar"}, []interface{}{1}, testAdap)

	if exec.query != `INSERT INTO "Foo" ("bar") VALUES (?);` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 1 || exec.args[0] != 1 {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestInsertMultiVal(t *testing.T) {
	exec := &execCap{}

	Insert(exec, "Foo", []string{"bar", "baz"}, []interface{}{1, true}, testAdap)

	if exec.query != `INSERT INTO "Foo" ("bar", "baz") VALUES (?, ?);` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 2 || exec.args[0] != 1 || exec.args[1] != true {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestDeleteWrongColsCount(t *testing.T) {
	idx := &testUniqueIdx{
		cols: []string{"bar"},
	}

	err := Delete(nil, "Foo", idx, testAdap)
	if err != ErrColsValsMismatch {
		t.Fatal("Expected error, but was", err)
	}
}

func TestDeleteNoCols(t *testing.T) {
	idx := &testUniqueIdx{}

	err := Delete(nil, "Foo", idx, testAdap)
	if err != ErrNoCols {
		t.Fatal("Expected error, but was", err)
	}
}

func TestDeleteOneCol(t *testing.T) {
	idx := &testUniqueIdx{
		cols: []string{"bar"},
		vals: []interface{}{1},
	}
	exec := &execCap{}
	if err := Delete(exec, "Foo", idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `DELETE FROM "Foo" WHERE "bar" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 1 || exec.args[0] != 1 {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestDeleteMultiCols(t *testing.T) {
	idx := &testUniqueIdx{
		cols: []string{"bar", "baz"},
		vals: []interface{}{1, true},
	}
	exec := &execCap{}
	if err := Delete(exec, "Foo", idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `DELETE FROM "Foo" WHERE "bar" = ? AND "baz" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 2 || exec.args[0] != 1 || exec.args[1] != true {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestUpdateWrongColCount(t *testing.T) {
	err := Update(nil, "Foo", []string{"bar"}, nil /*vals*/, nil, testAdap)
	if err != ErrColsValsMismatch {
		t.Fatal("Expected error, but was", err)
	}
}

func TestUpdateNoCols(t *testing.T) {
	err := Update(nil, "Foo", nil /*cols*/, nil /*vals*/, nil, testAdap)
	if err != ErrNoCols {
		t.Fatal("Expected error, but was", err)
	}
}

func TestUpdateWrongIdxColCount(t *testing.T) {
	cols := []string{"foo"}
	vals := []interface{}{1}
	idx := &testUniqueIdx{
		cols: []string{"bar"},
	}
	err := Update(nil, "Foo", cols, vals, idx, testAdap)
	if err != ErrColsValsMismatch {
		t.Fatal("Expected error, but was", err)
	}
}

func TestUpdateNoIdxCols(t *testing.T) {
	cols := []string{"foo"}
	vals := []interface{}{1}
	idx := &testUniqueIdx{}
	err := Update(nil, "Foo", cols, vals, idx, testAdap)
	if err != ErrNoCols {
		t.Fatal("Expected error, but was", err)
	}
}

func TestUpdateOneColOneIdxCol(t *testing.T) {
	cols := []string{"foo"}
	vals := []interface{}{1}
	idx := &testUniqueIdx{
		cols: []string{"bar"},
		vals: []interface{}{2},
	}
	exec := &execCap{}
	if err := Update(exec, "Foo", cols, vals, idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `UPDATE "Foo" SET "foo" = ? WHERE "bar" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 2 || exec.args[0] != 1 || exec.args[1] != 2 {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestUpdateOneColMultiIdxCol(t *testing.T) {
	cols := []string{"foo"}
	vals := []interface{}{1}
	idx := &testUniqueIdx{
		cols: []string{"bar", "baz"},
		vals: []interface{}{2, false},
	}
	exec := &execCap{}
	if err := Update(exec, "Foo", cols, vals, idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `UPDATE "Foo" SET "foo" = ? WHERE "bar" = ? AND "baz" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 3 || exec.args[0] != 1 || exec.args[1] != 2 || exec.args[2] != false {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestUpdateMultiColOneIdxCol(t *testing.T) {
	cols := []string{"foo", "bar"}
	vals := []interface{}{1, true}
	idx := &testUniqueIdx{
		cols: []string{"baz"},
		vals: []interface{}{2},
	}
	exec := &execCap{}
	if err := Update(exec, "Foo", cols, vals, idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `UPDATE "Foo" SET "foo" = ?, "bar" = ? WHERE "baz" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 3 || exec.args[0] != 1 || exec.args[1] != true || exec.args[2] != 2 {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestUpdateMultiColMultiIdxCol(t *testing.T) {
	cols := []string{"foo", "bar"}
	vals := []interface{}{1, true}
	idx := &testUniqueIdx{
		cols: []string{"baz", "qux"},
		vals: []interface{}{2, false},
	}
	exec := &execCap{}
	if err := Update(exec, "Foo", cols, vals, idx, testAdap); err != nil {
		t.Log(err)
		t.Fail()
	}
	if exec.query != `UPDATE "Foo" SET "foo" = ?, "bar" = ? WHERE "baz" = ? AND "qux" = ?;` {
		t.Log("Query didn't match", exec.query)
		t.Fail()
	}
	if len(exec.args) != 4 ||
		exec.args[0] != 1 || exec.args[1] != true || exec.args[2] != 2 || exec.args[3] != false {
		t.Log("Args didn't match", exec.args)
		t.Fail()
	}
}

func TestQuoteIdentifier(t *testing.T) {
	quoted := testAdap.Quote("foo")

	if quoted != `"foo"` {
		t.Fatal("not quoted", quoted)
	}
}

func TestQuoteIdentifierPanicsOnQuote(t *testing.T) {
	defer func() {
		val := recover()
		if val == nil {
			t.Fatal("expected a panic")
		}
	}()
	testAdap.Quote("f\"oo")

	t.Fatal("should never reach here")
}

func TestQuoteIdentifierPanicsOnNull(t *testing.T) {
	defer func() {
		val := recover()
		if val == nil {
			t.Fatal("expected a panic")
		}
	}()
	testAdap.Quote("f\x00oo")

	t.Fatal("should never reach here")
}

func TestAppendLock(t *testing.T) {

	expected := []struct {
		query string
		lock  Lock
	}{{"", LockNone}, {" FOR SHARE", LockRead}, {" FOR UPDATE", LockWrite}}
	for _, tuple := range expected {
		s := scanStmt{
			buf: new(bytes.Buffer),
			opts: Opts{
				Lock: tuple.lock,
			},
			adap: testAdap,
		}
		s.appendLock()
		newQuery := s.buf.String()
		if newQuery != tuple.query {
			t.Logf("Mismatched query %s != %s", newQuery, tuple.query)
			t.Fail()
		}
	}
}

func TestAppendLockPanicsOnBad(t *testing.T) {
	defer func() {
		val := recover()
		if val == nil {
			t.Fatal("expected a panic")
		}
	}()
	s := scanStmt{
		buf:  new(bytes.Buffer),
		name: "foo",
		opts: Opts{
			Lock: 3,
		},
		adap: testAdap,
	}
	s.appendLock()

	t.Fatal("should never reach here")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
