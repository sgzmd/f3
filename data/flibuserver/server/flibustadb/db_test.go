package flibustadb

import (
	"database/sql"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"sort"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	tok "github.com/liuzl/tokenizer"
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

	flibustaDb := NewFlibustaSqlite(db)

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

	flibustaDb := NewFlibustaSqlite(db)

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

	flibustaDb := NewFlibustaSqlite(db)
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

	flibustaDb := NewFlibustaSqlite(db)
	name, err := flibustaDb.GetSequenceName(34145)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Унесенный ветром", name)
}

func TestTokenizerForSQL(t *testing.T) {
	str := "Россия, которую мы!"

	words := tok.Tokenize(str)
	alnum, _ := regexp.Compile(`\p{L}+`)
	for _, w := range words {
		// Check if w is alphanumeric only
		if alnum.MatchString(w) {
			log.Printf("word: %+v", w)
		}
	}
	log.Printf("words: %+v", words)
}

// Tests MakeBooleanQuery
func TestMakeBooleanQuery(t *testing.T) {
	query := makeBooleanQuery("Россия, которую мы!")
	assert.Equal(t, "+россия +которую +мы", query)
}
