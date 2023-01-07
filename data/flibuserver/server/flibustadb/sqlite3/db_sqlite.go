package sqlite3

import (
	"database/sql"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
	"sync"
)

const (
	FLIBUSTA_DB = "../../../../../testutils/flibusta-test.db"
)

// Implements FlibustaDb interface for sqlite3 database.
type Sqlite3Database struct {
	sqliteDb        *sql.DB
	authorStatement *sql.Stmt
	seriesStatement *sql.Stmt
	lock            sync.Mutex
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

// GetSeriesBooks queries sqlite3 database for books in series.
func (s *Sqlite3Database) GetSeriesBooks(seriesId int64) ([]*pb.Book, error) {
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
