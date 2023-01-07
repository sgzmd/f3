package sqlite3

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// Tests for GetAuthorBooks
func TestGetAuthorBooks(t *testing.T) {
	db, err := sql.Open("sqlite3", FLIBUSTA_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	flibustaDb := NewSqlite3Db(db)

	books, err := flibustaDb.GetAuthorBooks(109170)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, books, 8)
	assert.Equal(t, "Чужие маски", books[0].BookName)
}
