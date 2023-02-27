package flibustadb

import (
	"database/sql"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	FLIBUSTA_DB = "../../../../testutils/flibusta-test.db"
)

// Tests for GetAuthorBooks
func TestGetAuthorBooks(t *testing.T) {
	db, err := sql.Open("sqlite3", FLIBUSTA_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	flibustaDb := NewFlibustaSqlDb(db)

	books, err := flibustaDb.GetAuthorBooks(109170)
	if err != nil {
		t.Fatal(err)
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].BookName < books[j].BookName
	})

	assert.Len(t, books, 8)
	assert.Equal(t, "Маска зверя", books[0].BookName)
}

// Tests for GetSeriesBooks
func TestGetSeriesBooks(t *testing.T) {
	db, err := sql.Open("sqlite3", FLIBUSTA_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	flibustaDb := NewFlibustaSqlDb(db)

	books, err := flibustaDb.GetSeriesBooks(34145)
	if err != nil {
		t.Fatal(err)
	}

	// sort books by BookName in-place
	sort.Slice(books, func(i, j int) bool {
		return books[i].BookName < books[j].BookName
	})

	assert.Len(t, books, 8)
	assert.Equal(t, "Маска зверя", books[0].BookName)
}

func TestGetAuthorName(t *testing.T) {
	db, err := sql.Open("sqlite3", FLIBUSTA_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	flibustaDb := NewFlibustaSqlDb(db)
	name, err := flibustaDb.GetAuthorName(109170)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, pb.AuthorName{
		LastName:   "Метельский",
		MiddleName: "Александрович",
		FirstName:  "Николай"}, name)
}

// Tests for GetSequenceName
func TestGetSequenceName(t *testing.T) {
	db, err := sql.Open("sqlite3", FLIBUSTA_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	flibustaDb := NewFlibustaSqlDb(db)
	name, err := flibustaDb.GetSequenceName(34145)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Унесенный ветром", name)
}
