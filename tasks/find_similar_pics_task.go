package tasks

import (
	"database/sql"
	"encoding/binary"

	"pixur.org/pixur/schema"
	"pixur.org/pixur/schema/db"
	tab "pixur.org/pixur/schema/tables"
	"pixur.org/pixur/status"
)

// TODO: add tests

type FindSimilarPicsTask struct {
	// Deps
	DB *sql.DB

	// Inputs
	PicID int64

	// Results
	SimilarPicIDs []int64
}

func (t *FindSimilarPicsTask) Run() (errCap error) {
	t.SimilarPicIDs = make([]int64, 0) // set default, to make json encoding easier
	j, err := tab.NewJob(t.DB)
	if err != nil {
		return status.InternalError(err, "can't create new job")
	}
	defer cleanUp(j, errCap)

	pics, err := j.FindPics(db.Opts{
		Start: tab.PicsPrimary{&t.PicID},
		Limit: 1,
	})
	if err != nil {
		return status.InternalError(err, "can't lookup pic")
	}
	if len(pics) != 1 {
		return status.InvalidArgument(err, "can't lookup pic", len(pics))
	}
	pic := pics[0]

	dctIdentType := schema.PicIdent_DCT_0

	picIdents, err := j.FindPicIdents(db.Opts{
		Start: tab.PicIdentsPrimary{PicId: &t.PicID, Type: &dctIdentType},
		Limit: 1,
	})
	if err != nil {
		return status.InternalError(err, "can't lookup pic ident")
	}
	if len(picIdents) != 1 {
		return status.InvalidArgument(err, "can't lookup pic ident", len(picIdents))
	}
	targetIdent := picIdents[0]
	match := binary.BigEndian.Uint64(targetIdent.Value)

	scanOpts := db.Opts{
		Start: tab.PicIdentsIdent{Type: &dctIdentType},
	}
	err = j.ScanPicIdents(scanOpts, func(pi *schema.PicIdent) error {
		if pi.PicId == pic.PicId {
			return nil
		}
		guess := binary.BigEndian.Uint64(pi.Value)
		bits := guess ^ match
		bitCount := 0
		// replace this with something that isn't hideously slow.  Hamming distance would be
		// better served by a look up table or some 64 bit specific bit magic.  Cosine similarity
		// on the attached floats would also work.
		for i := uint(0); i < 64; i++ {
			if ((1 << i) & bits) > 0 {
				bitCount++
			}
		}
		if bitCount <= 10 {
			t.SimilarPicIDs = append(t.SimilarPicIDs, pi.PicId)
		}

		return nil
	})

	if err != nil {
		return status.InternalError(err, "can't scan pic idents")
	}

	return nil
}
