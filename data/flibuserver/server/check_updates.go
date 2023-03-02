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

	// We will start with a very naive and simple implementation
	for _, entry := range in.TrackedEntry {
		var statement *sql.Stmt
		key := entry.Key

		// An artifact of the old code - previously used direct SQL statements, using
		// methods instead of SQL statements now.
		var stm func(_ int64) ([]*proto.Book, error)
		if key.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
			stm = s.db.GetAuthorBooks
		} else if key.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
			stm = s.db.GetSeriesBooks
		}

		newBooks, err := stm(key.EntityId)
		if err != nil {
			return nil, err
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
