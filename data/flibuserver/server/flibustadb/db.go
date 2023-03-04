package flibustadb

import (
	"database/sql"
	"fmt"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb/mariadb"
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

// enum with fields MARIA_DB and SQLITE
type DbEngine int

const (
	MARIA_DB DbEngine = iota
	SQLITE   DbEngine = iota
)

// FlibustaDbSql Implements FlibustaDb interface for SQL database.
type FlibustaDbSql struct {
	mariaDb *sql.DB
	lock    sync.Mutex

	engine DbEngine
}

// GetAuthorBooks returns all books by author.
func (s *FlibustaDbSql) GetAuthorBooks(authorId int64) ([]*pb.Book, error) {
	const query = `select b.BookId, b.Title, a.Pos from libbook b, libavtor a 
					   where b.BookId = a.BookId and a.AvtorId = ? and b.Deleted != '1'`

	rows, err := s.mariaDb.Query(query, authorId)
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
	const query = `
				select an.FirstName, an.MiddleName, an.LastName 
				from libavtorname an
				where an.AvtorId = ?`

	rs, err := s.mariaDb.Query(query, authorid)
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

// GetSequenceName implements FlibustaDb.GetSequenceName for MariaDB
func (s *FlibustaDbSql) GetSequenceName(seqId int64) (string, error) {
	rs, err := s.mariaDb.Query(`
				select s.SeqName 
				from libseqname s
				where s.SeqId = ?`, seqId)
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
	rows, err := s.mariaDb.Query(`select b.BookId, b.Title, s.SeqNumb from libbook b, libseq s 
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
	rows, err := s.mariaDb.Query(`select AvtorId from libavtor where BookId = ?`, bookId)
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
	return s.mariaDb.Close()
}

// NewFlibustaSqlite Creates new FlibustaDbSql instance for SQLite3 database.
func NewFlibustaSqlite(db *sql.DB) *FlibustaDbSql {
	return &FlibustaDbSql{mariaDb: db, engine: SQLITE}
}

// NewFlibustaSqlMariaDb Creates new FlibustaDbSql instance for MariaDB database.
func NewFlibustaSqlMariaDb(db *sql.DB) *FlibustaDbSql {
	return &FlibustaDbSql{mariaDb: db, engine: MARIA_DB}
}

// SearchAuthors searches authors by name.
func (s *FlibustaDbSql) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	var query string
	if s.engine == SQLITE {
		query = fmt.Sprintf(sqlite3.AuthorQueryTemplateSqlite, req.SearchTerm)
	} else {
		query = fmt.Sprintf(mariadb.SearchAuthorsFtsMysql, req.SearchTerm)
	}
	log.Printf("SearchAuthors: sql=%s", query)
	rows, err := s.mariaDb.Query(query)
	if err != nil {
		return nil, err
	}

	var entries = make([]*pb.FoundEntry, 0, 10)
	for rows.Next() {
		entry := &pb.FoundEntry{EntryType: pb.EntryType_ENTRY_TYPE_AUTHOR}

		err = rows.Scan(&entry.Author, &entry.EntryId, &entry.NumEntities)
		if err != nil {
			log.Fatalf("Failed to scan the row: %v", err)
			return nil, err
		}
		entry.EntryName = entry.Author

		log.Printf("SearchAuthors: entry=%v", entry)

		entries = append(entries, entry)
	}

	return entries, nil
}

// SearchSeries searches series by name.
func (s *FlibustaDbSql) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	var query string
	if s.engine == SQLITE {
		query = fmt.Sprintf(sqlite3.SequenceQueryTemplateSqlite, req.SearchTerm)
	} else {
		query = fmt.Sprintf(mariadb.SearchSeriesFtsMysql, req.SearchTerm)
	}
	log.Printf("SearchSeries: sql=%s", query)
	rows, err := s.mariaDb.Query(query)
	if err != nil {
		return nil, err
	}

	// iterate over rows
	var entries = make([]*pb.FoundEntry, 0, 10)
	for rows.Next() {
		entry := &pb.FoundEntry{EntryType: pb.EntryType_ENTRY_TYPE_SERIES}

		err = rows.Scan(&entry.EntryId, &entry.EntryName, &entry.Author, &entry.NumEntities)
		if err != nil {
			log.Fatalf("Failed to scan the row: %v", err)
			return nil, err
		}

		log.Printf("SearchSeries: entry=%v", entry)

		entries = append(entries, entry)
	}

	return entries, nil
}
