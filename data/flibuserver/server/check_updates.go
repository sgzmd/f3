package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
)

// CheckUpdates Searches for updates in the collection of tracked entries.
// Implementation is very straightforward and not very performant
// but it's possible that it's good enough.
// See: ../proto/flibustier.proto for proto definitions.
func (s *server) CheckUpdates(_ context.Context, in *proto.CheckUpdatesRequest) (*proto.CheckUpdatesResponse, error) {
	log.Printf("Received: %v", in)

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	response := make([]*proto.UpdateRequired, 0)

	astm, err := s.sqliteDb.Prepare(`
		select b.BookId, b.Title from libbook b, libavtor a 
		where b.BookId = a.BookId and a.AvtorId = ? and b.Deleted != '1'
`)

	if err != nil {
		return nil, err
	}

	sstm, err := s.sqliteDb.Prepare(`
	 
	`)
	if err != nil {
		return nil, err
	}

	// We will start with a very naive and simple implementation
	for _, entry := range in.TrackedEntry {
		var rs *sql.Rows
		var err error

		var statement *sql.Stmt
		key := entry.Key
		if key.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
			statement = astm
		} else if key.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
			statement = sstm
		}

		rs, err = statement.Query(key.EntityId)
		if err != nil {
			return nil, err
		}

		newBooks := make([]*proto.Book, 0)

		for rs.Next() {
			var bookId int32
			var title string
			rs.Scan(&bookId, &title)

			newBooks = append(newBooks, &proto.Book{BookName: title, BookId: bookId})
		}

		if len(newBooks) == 0 {
			return nil,
				fmt.Errorf("exceptional situation: nothing was found for %v with EntryId %d",
					statement,
					key.EntityId)
		}

		if entry.NumEntries != int32(len(newBooks)) {
			// If it is equal, no updates required

			// sort.SliceStable(entry.Book, CreateBookComparator(entry.Book))
			// sort.SliceStable(new_books, CreateBookComparator(new_books))
			oldBookMap := make(map[int]*proto.Book)
			for _, b := range entry.Book {
				oldBookMap[int(b.BookId)] = b
			}

			if len(newBooks) <= int(entry.NumEntries) {
				continue
			}

			newlyAddedBooks := make([]*proto.Book, 0, len(newBooks)-int(entry.NumEntries))
			for _, b := range newBooks {
				_, exists := oldBookMap[int(b.BookId)]
				if !exists {
					// Well we found the missing book
					newlyAddedBooks = append(newlyAddedBooks, b)
				}
			}

			if len(newlyAddedBooks) > 0 {
				response = append(response, &proto.UpdateRequired{
					TrackedEntry:  entry,
					NewNumEntries: int32(len(newBooks)),
					NewBook:       newlyAddedBooks,
				})
			}
		}
	}

	return &proto.CheckUpdatesResponse{UpdateRequired: response}, nil
}
