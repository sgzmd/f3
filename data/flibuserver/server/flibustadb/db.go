package flibustadb

import (
	"database/sql"
	"fmt"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb/sqlite3"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
	"sync"
)

type FlibustaDb interface {
	SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error)
	SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error)
	GetAuthorBooks(authorId int64) ([]*pb.Book, error)
	GetSeriesBooks(seriesId int64) ([]*pb.Book, error)

	GetAuthorName(authorId int64) (pb.AuthorName, error)
	GetSequenceName(seqId int64) (string, error)

	GetBookAuthor(bookId int64) (pb.AuthorName, error)

	Close() error
}

// FlibustaDbSql Implements FlibustaDb interface for SQL database.
type FlibustaDbSql struct {
	db                    *sql.DB
	db2                   *sql.DB
	authorStatement       *sql.Stmt
	seriesStatement       *sql.Stmt
	authorNameStatement   *sql.Stmt
	sequenceNameStatement *sql.Stmt

	lock sync.Mutex
}

// GetAuthorBooks returns all books by author.
func (s *FlibustaDbSql) GetAuthorBooks(authorId int64) ([]*pb.Book, error) {
	{
		// Scoped lock which will be unlocked whether we create a new statement or not.
		s.lock.Lock()
		defer s.lock.Unlock()
		var err error
		if s.authorStatement == nil {
			s.authorStatement, err = s.db.Prepare(
				`select b.BookId, b.Title, a.Pos from libbook b, libavtor a 
					   where b.BookId = a.BookId and a.AvtorId = ? and b.Deleted != '1'`)
			if err != nil {
				return nil, err
			}
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
		var pos int32
		err := rows.Scan(&bookId, &bookTitle, &pos)

		if err != nil {
			return nil, err
		}

		books = append(books, &pb.Book{BookId: bookId, BookName: bookTitle, OrderInSequence: pos})
	}

	return books, nil
}

func (s *FlibustaDbSql) GetAuthorName(authorid int64) (pb.AuthorName, error) {
	{
		s.lock.Lock()
		defer s.lock.Unlock()

		var err error
		if s.authorNameStatement == nil {
			s.authorNameStatement, err = s.db.Prepare(`
				select an.FirstName, an.MiddleName, an.LastName 
				from libavtorname an
				where an.AvtorId = ?`)
			if err != nil {
				return pb.AuthorName{}, err
			}
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
func (s *FlibustaDbSql) GetSequenceName(seqId int64) (string, error) {
	{
		s.lock.Lock()
		defer s.lock.Unlock()

		if s.sequenceNameStatement == nil {
			var err error
			s.sequenceNameStatement, err = s.db.Prepare(`
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
func (s *FlibustaDbSql) GetSeriesBooks(seriesId int64) ([]*pb.Book, error) {
	db := s.db
	if s.db2 != nil {
		db = s.db2
	}

	//{
	//	// Scoped lock which will be unlocked whether we create a new statement or not.
	//	s.lock.Lock()
	//	defer s.lock.Unlock()
	//	if s.seriesStatement == nil {
	//		s.seriesStatement, _ = db.Prepare(
	//			`select b.BookId, b.Title, s.SeqNumb from libbook b, libseq s
	//				   where s.BookId = b.BookId and s.SeqId = ? and b.Deleted != '1'
	//				   order by s.SeqNumb`)
	//	}
	//}

	rows, err := db.Query(`select b.BookId, b.Title, s.SeqNumb from libbook b, libseq s 
					   where s.BookId = b.BookId and s.SeqId = ? and b.Deleted != '1' 
					   order by s.SeqNumb`, seriesId)

	if err != nil {
		return nil, err
	}

	log.Printf("GetSeriesBooks(%d)", seriesId)

	books := make([]*pb.Book, 0)
	for rows.Next() {
		var bookId int32
		var bookTitle string
		var seqNumb int32
		err := rows.Scan(&bookId, &bookTitle, &seqNumb)

		log.Printf("GetSeriesBooks(%d) bookId=%d, bookTitle=%s, seqNumb=%d", seriesId, bookId, bookTitle, seqNumb)

		if err != nil {
			return nil, err
		}

		books = append(books, &pb.Book{BookId: bookId, BookName: bookTitle, OrderInSequence: seqNumb})
	}

	log.Printf("GetSeriesBooks(%d) total books=%d, books=%v", seriesId, len(books), books)

	return books, nil
}

// GetBookAuthor implements FlibustaDb.GetBookAuthor for Sqlite3.
func (s *FlibustaDbSql) GetBookAuthor(bookId int64) (pb.AuthorName, error) {
	rows, err := s.db.Query(`select AvtorId from libavtor where BookId = ?`, bookId)
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
func (s *FlibustaDbSql) Close() error {
	return s.db.Close()
}

// NewFlibustaSqlDb creates new FlibustaDbSql instance.
func NewFlibustaSqlDb(sqliteDb *sql.DB) *FlibustaDbSql {
	return &FlibustaDbSql{db: sqliteDb, db2: nil, authorStatement: nil, seriesStatement: nil}
}

func NewFlibustaSqlDbWithMaria(sqliteDb *sql.DB, mariaDb *sql.DB) *FlibustaDbSql {
	return &FlibustaDbSql{db: sqliteDb, db2: mariaDb, authorStatement: nil, seriesStatement: nil}
}

func (s *FlibustaDbSql) iterateOverAuthors(query string) ([]*pb.FoundEntry, error) {
	rows, err := s.db.Query(query)

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

func (s *FlibustaDbSql) iterateOverSeries(query string) ([]*pb.FoundEntry, error) {
	rows, err := s.db.Query(query)

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

// SearchAuthors searches authors by name. Specific for SQLite3 implementation.
func (s *FlibustaDbSql) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	query := sqlite3.CreateAuthorSearchQuery(req.SearchTerm)
	return s.iterateOverAuthors(query)
}

// SearchSeries searches series by name. Specific for SQLite3 implementation.
func (s *FlibustaDbSql) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	query := sqlite3.CreateSequenceSearchQuery(req.SearchTerm)
	return s.iterateOverSeries(query)
}
