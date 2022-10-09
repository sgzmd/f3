package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb/sqlite3"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
	"strings"
	"sync"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/golang/protobuf/proto"

	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	pb.UnimplementedFlibustierServiceServer
	sqliteDb *sql.DB
	data     *badger.DB
	Lock     sync.RWMutex
}

const (
	TrackedEntryPrefix = "tracked_entry_"
	UserEntryPrefix    = "user_entry_"
)

func (s *server) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	log.Printf("Searching for author: %s", req)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	// TODO: Refactor this part so we can re-use iteration code with a different SQL
	// query - as long as it supplies the right results.
	query := sqlite3.CreateAuthorSearchQuery(req.SearchTerm)
	return s.iterateOverAuthors(query)
}

func (s *server) iterateOverAuthors(query string) ([]*pb.FoundEntry, error) {
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

func (s *server) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	log.Printf("Searching for series: %s", req)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	query := sqlite3.CreateSequenceSearchQuery(req.SearchTerm)
	return s.iterateOverSeries(query)
}

func (s *server) getBooks(query string) ([]*pb.Book, error) {
	rows, err := s.sqliteDb.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	books := make([]*pb.Book, 0, 10)
	for rows.Next() {
		var bookName string
		var bookId int32

		err = rows.Scan(&bookName, &bookId)
		if err != nil {
			return nil, err
		}

		books = append(books, &pb.Book{
			BookName: bookName,
			BookId:   bookId,
		})
	}

	return books, nil
}

func (s *server) iterateOverSeries(query string) ([]*pb.FoundEntry, error) {
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

func (s *server) GlobalSearch(_ context.Context, in *pb.GlobalSearchRequest) (*pb.GlobalSearchResponse, error) {
	log.Printf("Received: %v", in.GetSearchTerm())

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	var entries []*pb.FoundEntry = make([]*pb.FoundEntry, 0, 10)

	// If there's no filter for series
	if in.EntryTypeFilter != pb.EntryType_ENTRY_TYPE_SERIES {
		authors, err := s.SearchAuthors(in)
		if err != nil {
			return nil, err
		}
		entries = append(entries, authors...)
	}

	if in.EntryTypeFilter != pb.EntryType_ENTRY_TYPE_AUTHOR {
		series, err := s.SearchSeries(in)
		if err != nil {
			return nil, err
		}
		entries = append(entries, series...)
	}

	return &pb.GlobalSearchResponse{
		OriginalRequest: in,
		Entry:           entries,
	}, nil
}

func GetEntityBooks(sql *sql.Stmt, entityId int32) ([]*pb.Book, error) {
	rs, err := sql.Query(entityId)

	if err != nil {
		return nil, err
	}

	books := make([]*pb.Book, 0)
	for rs.Next() {
		var bookTitle string
		var bookId int32

		rs.Scan(&bookTitle, &bookId)
		books = append(books, &pb.Book{BookId: bookId, BookName: bookTitle})
	}

	return books, nil
}

func (s *server) GetAuthorBooks(_ context.Context, in *pb.GetAuthorBooksRequest) (*pb.GetAuthorBooksResponse, error) {
	log.Printf("GetAuthorBooks: %+v", in)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	sql, err := s.sqliteDb.Prepare(`
		select 
		  lb.Title,
		  lb.Bookid
		from libbook lb, libavtor la, author_fts a
		where la.BookId = lb.BookId 
		and a.authorId = la.AvtorId
		and lb.Deleted != '1'
		and la.AvtorId = ?
		group by la.BookId order by la.BookId;`)

	if err != nil {
		return nil, err
	}
	books, err := GetEntityBooks(sql, in.AuthorId)

	if err != nil {
		return nil, err
	}

	sql, err = s.sqliteDb.Prepare(`
		select an.FirstName, an.MiddleName, an.LastName 
		from libavtorname an
		where an.AvtorId = ?`)
	if err != nil {
		return nil, err
	}
	rs, err := sql.Query(in.AuthorId)
	if err != nil {
		return nil, err
	}

	if rs.Next() {
		var firstName, middleName, lastName string
		rs.Scan(&firstName, &middleName, &lastName)
		name := &pb.EntityName{Name: &pb.EntityName_AuthorName{
			AuthorName: &pb.AuthorName{
				FirstName:  firstName,
				MiddleName: middleName,
				LastName:   lastName}}}

		return &pb.GetAuthorBooksResponse{
			EntityBookResponse: &pb.EntityBookResponse{
				Book: books, EntityId: in.AuthorId, EntityName: name}}, nil
	}

	return nil, fmt.Errorf("no author associated with id %d", in.AuthorId)
}

func (s *server) GetSeriesBooks(ctx context.Context, in *pb.GetSeriesBooksRequest) (*pb.GetSeriesBooksResponse, error) {
	log.Printf("GetSeriesBooks: %+v", in)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	sql, err := s.sqliteDb.Prepare(`
		SELECT b.Title, b.BookId
		FROM libseq ls, libseqname lsn , libbook b
		WHERE ls.seqId = lsn.seqId and ls.seqId = ? and ls.BookId = b.BookId and b.Deleted != '1'
				  group by b.BookId
				  order by ls.SeqNumb;`)

	if err != nil {
		return nil, err
	}
	books, err := GetEntityBooks(sql, in.SequenceId)

	if err != nil {
		return nil, err
	}

	rs, err := s.sqliteDb.Query("select SeqName from libseqname where SeqId = ?", in.SequenceId)
	if err != nil {
		return nil, err
	}
	if rs.Next() {
		var seqName string
		rs.Scan(&seqName)
		name := &pb.EntityName{Name: &pb.EntityName_SequenceName{SequenceName: seqName}}

		return &pb.GetSeriesBooksResponse{
			EntityBookResponse: &pb.EntityBookResponse{
				Book: books, EntityId: in.SequenceId, EntityName: name}}, nil
	}

	return nil, fmt.Errorf("no series associated with id %d", in.SequenceId)
}

func (s *server) GetUserInfo(_ context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	ui := pb.UserInfo{}
	prefix := []byte((UserEntryPrefix + in.UserId))

	err := s.data.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			skey := strings.TrimPrefix(string(it.Item().Key()), UserEntryPrefix)
			if skey == in.UserId {
				return it.Item().Value(func(val []byte) error {
					err := proto.Unmarshal(val, &ui)
					if err != nil {
						return err
					}
					return nil
				})
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	} else {
		if ui.UserId == "" {
			if in.GetAction() == pb.UserInfoAction_USER_INFO_ACTION_CREATE {
				ui.UserId = in.UserId
				ui.UserTelegramId = in.UserTelegramId

				err := s.data.Update(func(txn *badger.Txn) error {
					val, err := proto.Marshal(&ui)
					if err != nil {
						return err
					}

					txn.Set(prefix, val)

					return nil
				})

				if err != nil {
					return nil, err
				} else {
					return &pb.GetUserInfoResponse{UserInfo: &ui}, nil
				}
			}
		} else {
			return &pb.GetUserInfoResponse{UserInfo: &ui}, nil
		}

		return nil, errors.New("User not found")
	}
}

func (s *server) ListUsers(_ context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	prefix := []byte((UserEntryPrefix))
	resp := pb.ListUsersResponse{
		User: make([]*pb.UserInfo, 0, 10),
	}
	err := s.data.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			err := it.Item().Value(func(val []byte) error {
				ui := pb.UserInfo{}
				err := proto.Unmarshal(val, &ui)
				if err != nil {
					return err
				}
				resp.User = append(resp.User, &ui)
				return nil
			})

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	} else {
		return &resp, nil
	}
}

// DeleteAllUsers deletes all users - for testing only.
func (s *server) DeleteAllUsers(_ context.Context, _ *pb.DeleteAllUsersRequest) (*pb.DeleteAllUsersResponse, error) {
	e := s.data.DropPrefix([]byte(UserEntryPrefix))
	if e != nil {
		return nil, e
	} else {
		return &pb.DeleteAllUsersResponse{}, nil
	}
}

func (s *server) DeleteAllTracked(_ context.Context, _ *pb.DeleteAllTrackedRequest) (*pb.DeleteAllTrackedResponse, error) {
	e := s.data.DropPrefix([]byte(TrackedEntryPrefix))
	if e != nil {
		return nil, e
	} else {
		return &pb.DeleteAllTrackedResponse{}, nil
	}
}

func (s *server) Close() {
	log.Println("Closing database connection.")
	s.sqliteDb.Close()
}

func (s *server) Shutdown() {
	s.sqliteDb.Close()
	s.data.Close()
}
