package sqlite3

import (
	"database/sql"
	"fmt"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
	"sync"
)

const (
	FLIBUSTA_DB = "../../../../../testutils/flibusta-test.db"
)

// Sqlite3Database Implements FlibustaDb interface for sqlite3 database.
type Sqlite3Database struct {
	db                    *flibustadb.FlibustaDb
	sqliteDb              *sql.DB
	authorStatement       *sql.Stmt
	seriesStatement       *sql.Stmt
	authorNameStatement   *sql.Stmt
	sequenceNameStatement *sql.Stmt

	lock sync.Mutex
}

// SearchAuthors searches authors by name.
func (s *Sqlite3Database) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	query := CreateAuthorSearchQuery(req.SearchTerm)
	return s.iterateOverAuthors(query)
}

// SearchSeries searches series by name.
func (s *Sqlite3Database) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	query := CreateSequenceSearchQuery(req.SearchTerm)
	return s.iterateOverSeries(query)
}

// GetAuthorBooks returns all books by author.
func (s *Sqlite3Database) GetAuthorBooks(authorId int64) ([]*pb.Book, error) {
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

func (s *Sqlite3Database) GetAuthorName(authorid int64) (pb.AuthorName, error) {
	{
		s.lock.Lock()
		defer s.lock.Unlock()

		if s.authorNameStatement == nil {
			s.authorNameStatement, _ = s.sqliteDb.Prepare(`
				select an.FirstName, an.MiddleName, an.LastName 
				from libavtorname an
				where an.AvtorId = ?`)
		}
	}

	rs, err := s.authorNameStatement.Query(authorid)
	name := pb.AuthorName{}
	if err != nil {
		return name, err
	}

	if rs.Next() {
		rs.Scan(&name.FirstName, &name.MiddleName, &name.LastName)
		return name, nil
	} else {
		return name, fmt.Errorf("author with authorId=%d not found", authorid)
	}
}

// GetSequenceName implements FlibustaDb.GetSequenceName for Sqlite3
func (s *Sqlite3Database) GetSequenceName(seqId int64) (string, error) {
	{
		s.lock.Lock()
		defer s.lock.Unlock()

		if s.sequenceNameStatement == nil {
			var err error
			s.sequenceNameStatement, err = s.sqliteDb.Prepare(`
				select s.SeqName 
				from libseqname s
				where s.SeqId = ?`)

			if err != nil {
				return "", err
			}
		}
	}

	rs, err := s.sequenceNameStatement.Query(seqId)
	if err != nil {
		return "", err
	}

	if rs.Next() {
		var name string
		rs.Scan(&name)
		return name, nil
	} else {
		return "", fmt.Errorf("sequence with seqId=%d not found", seqId)
	}
}

// GetSeriesBooks queries sqlite3 database for books in series.
func (s *Sqlite3Database) GetSeriesBooks(seriesId int64) ([]*pb.Book, error) {
	{
		// Scoped lock which will be unlocked whether we create a new statement or not.
		s.lock.Lock()
		defer s.lock.Unlock()
		if s.seriesStatement == nil {
			s.seriesStatement, _ = s.sqliteDb.Prepare(
				`select b.BookId, b.Title from libbook b, libseq s 
					   where s.BookId = b.BookId and s.SeqId = ? and b.Deleted != '1' 
					   order by s.SeqNumb`)
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

// GetBookAuthor implements FlibustaDb.GetBookAuthor for Sqlite3.
func (s *Sqlite3Database) GetBookAuthor(bookId int64) (pb.AuthorName, error) {
	rows, err := s.sqliteDb.Query(`select AvtorId from libavtor where BookId = ?`, bookId)
	if err != nil {
		return pb.AuthorName{}, err
	}

	if rows.Next() {
		var authorId int64
		rows.Scan(&authorId)
		return s.GetAuthorName(authorId)
	} else {
		return pb.AuthorName{}, fmt.Errorf("author not found for bookId=%d", bookId)
	}
}

// Close implements FlibustaDb.Close for Sqlite3
func (s *Sqlite3Database) Close() error {
	return s.sqliteDb.Close()
}

// NewSqlite3Db creates new Sqlite3Database instance.
func NewSqlite3Db(sqliteDb *sql.DB) *Sqlite3Database {
	return &Sqlite3Database{sqliteDb: sqliteDb, authorStatement: nil, seriesStatement: nil}
}

func (s *Sqlite3Database) iterateOverAuthors(query string) ([]*pb.FoundEntry, error) {
	rows, err := s.sqliteDb.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*pb.FoundEntry = make([]*pb.FoundEntry, 0, 10)
	for rows.Next() {
		var authorName string
		var authorId int64
		var count int32

		err = rows.Scan(&authorName, &authorId, &count)
		if err != nil {
			log.Fatalf("Failed to scan the row: %v", err)
			return nil, err
		}

		entries = append(entries, &pb.FoundEntry{
			EntryType:   pb.EntryType_ENTRY_TYPE_AUTHOR,
			Author:      authorName,
			EntryName:   authorName,
			EntryId:     authorId,
			NumEntities: count,
		})
	}

	return entries, nil
}

func (s *Sqlite3Database) iterateOverSeries(query string) ([]*pb.FoundEntry, error) {
	rows, err := s.sqliteDb.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*pb.FoundEntry = make([]*pb.FoundEntry, 0, 10)
	for rows.Next() {
		var seqName string
		var authors string
		var seqId int64
		var count int32

		err = rows.Scan(&seqName, &authors, &seqId, &count)
		if err != nil {
			log.Fatalf("Failed to scan the row: %v", err)
			return nil, err
		}

		entries = append(entries, &pb.FoundEntry{
			EntryType:   pb.EntryType_ENTRY_TYPE_SERIES,
			Author:      authors,
			EntryName:   seqName,
			EntryId:     seqId,
			NumEntities: count,
		})
	}

	return entries, nil
}
