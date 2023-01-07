package sqlite3

import (
	"database/sql"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"sync"
)

// Implements FlibustaDb interface for sqlite3 database.
type Sqlite3Db struct {
	sqliteDb        *sql.DB
	authorStatement *sql.Stmt
	seriesStatement *sql.Stmt
	lock            sync.Mutex
}

// SearchAuthors searches authors by name.
func (s *Sqlite3Db) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	return nil, nil
}

// SearchSeries searches series by name.
func (s *Sqlite3Db) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	return nil, nil
}

// GetAuthorBooks returns all books by author.
func (s *Sqlite3Db) GetAuthorBooks(authorId int64) ([]*pb.Book, error) {
	{
		// Scoped lock which will be unlocked whether we create a new statement or not.
		s.lock.Lock()
		defer s.lock.Unlock()
		if s.authorStatement == nil {
			s.authorStatement, _ = s.sqliteDb.Prepare(
				`select b.BookId, b.Title from libbook b, libavtor a 
					   where b.BookId = a.BookId and a.AvtorId = ? and b.Deleted != '1'`)
		}
	}

	rows, err := s.authorStatement.Query(authorId)
	if err != nil {
		return nil, err
	}

	books := make([]*pb.Book, 0)
	for rows.Next() {
		var bookId int32
		var bookTitle string
		err := rows.Scan(&bookId, &bookTitle)

		if err != nil {
			return nil, err
		}

		books = append(books, &pb.Book{BookId: bookId, BookName: bookTitle})
	}

	return books, nil
}

// GetSeriesBooks queries sqlite3 database for books in series.
func (s *Sqlite3Db) GetSeriesBooks(seriesId int64) ([]*pb.Book, error) {
	{
		// Scoped lock which will be unlocked whether we create a new statement or not.
		s.lock.Lock()
		defer s.lock.Unlock()
		if s.seriesStatement == nil {
			s.seriesStatement, _ = s.sqliteDb.Prepare(
				`select b.BookId, b.Title from libbook b, libseq s 
					   where s.BookId = b.BookId and s.SeqId = ? and b.Deleted != '1'`)
		}
	}

	rows, err := s.seriesStatement.Query(seriesId)
	if err != nil {
		return nil, err
	}

	books := make([]*pb.Book, 0)
	for rows.Next() {
		var bookId int32
		var bookTitle string
		err := rows.Scan(&bookId, &bookTitle)

		if err != nil {
			return nil, err
		}

		books = append(books, &pb.Book{BookId: bookId, BookName: bookTitle})
	}

	return books, nil
}

// NewSqlite3Db creates new Sqlite3Db instance.
func NewSqlite3Db(sqliteDb *sql.DB) *Sqlite3Db {
	return &Sqlite3Db{sqliteDb: sqliteDb, authorStatement: nil, seriesStatement: nil}
}

// Compare this snippet from flibuserver\server\flibustadb\db_test.go:
// package flibustadb
//
// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestSqlite3Db(t *testing.T) {
// 	db := NewSqlite3Db()
// 	assert.NotNil(t, db)
// }
//
// Compare this snippet from flibuserver\server\flibustadb\badgerdb\badgerdb.go:
// package badgerdb
//
// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
//
// 	"github.com/dgraph-io/badger"
// 	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
// 	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1/ftypes"
// 	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1/ftypes/ftypespb"
// 	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1/ftypes/ftypespb/ftypespbpb"
// 	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1/ftypes/ftypespb/ftypespbpb/ftypespbpbpb"
