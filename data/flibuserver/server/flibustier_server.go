package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb"
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
	data *badger.DB
	Lock sync.RWMutex
	db   flibustadb.FlibustaDb
}

const (
	TrackedEntryPrefix = "tracked_entry_"
	UserEntryPrefix    = "user_entry_"
)

func (s *server) SearchAuthors(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	log.Printf("Searching for author: %s", req)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return s.db.SearchAuthors(req)
}

func (s *server) SearchSeries(req *pb.GlobalSearchRequest) ([]*pb.FoundEntry, error) {
	log.Printf("Searching for series: %s", req)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return s.db.SearchSeries(req)
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

	books, err := s.db.GetAuthorBooks(int64(in.AuthorId))
	if err != nil {
		return nil, err
	}

	authorName, err := s.db.GetAuthorName(int64(in.AuthorId))
	if err != nil {
		return nil, err
	}
	name := &pb.EntityName{Name: &pb.EntityName_AuthorName{AuthorName: &authorName}}

	return &pb.GetAuthorBooksResponse{
		EntityBookResponse: &pb.EntityBookResponse{
			Book: books, EntityId: in.AuthorId, EntityName: name}}, nil
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
