package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
	proto2 "github.com/golang/protobuf/proto"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *server) TrackEntry(_ context.Context, req *proto.TrackEntryRequest) (*proto.TrackEntryResponse, error) {
	log.Printf("TrackEntryRequest: %+v", req)

	key := req.Key
	if !(req.Key.EntityId > 0) {
		return nil, createRequestError(NoEntryId)
	}

	if req.Key.UserId == "" {
		return nil, createRequestError(NoUserId)
	}

	if req.Key.EntityType == proto.EntryType_ENTRY_TYPE_UNSPECIFIED {
		return nil, createRequestError(NoEntryType)
	}

	alreadyTracked := false
	err := s.data.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		entryKey := &proto.KvRecordKey{}
		entryKey.Key = &proto.KvRecordKey_TrackedEntryKey{TrackedEntryKey: key}

		prefix := TrackedEntryPrefix + proto2.MarshalTextString(entryKey)

		bprefix := []byte(prefix)

		for it.Seek(bprefix); it.ValidForPrefix(bprefix); it.Next() {
			alreadyTracked = true
			return nil
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if alreadyTracked && !req.ForceUpdate {
		return &proto.TrackEntryResponse{Key: key, Result: proto.TrackEntryResult_TRACK_ENTRY_RESULT_ALREADY_TRACKED}, nil
	}

	// So we will be definitely tracking this, let's obtain all the info
	// about this entry.
	books, err := s.GetBooksForKey(key)
	if err != nil {
		return nil, err
	}

	s.data.Update(func(txn *badger.Txn) error {
		entryKey := &proto.KvRecordKey{Key: &proto.KvRecordKey_TrackedEntryKey{TrackedEntryKey: key}}

		prefix := TrackedEntryPrefix + proto2.MarshalTextString(entryKey)
		keyBytes := []byte(prefix)

		if err != nil {
			return err
		}

		entryName, err := s.GetEntryName(key)
		if err != nil {
			return err
		}

		entryAuthor, err := s.GetEntryAuthor(key)
		if err != nil {
			return err
		}

		val := proto.TrackedEntry{
			Key:         key,
			EntryName:   entryName,
			NumEntries:  int32(len(books)),
			Book:        books,
			EntryAuthor: entryAuthor,
			Status:      req.Status,
		}

		now := time.Now()
		val.Saved = &timestamppb.Timestamp{Seconds: now.Unix()}

		value, err := proto2.Marshal(&val)
		if err != nil {
			return err
		}

		return txn.Set(keyBytes, value)
	})

	if err != nil {
		return nil, err
	}

	return &proto.TrackEntryResponse{Key: key, Result: proto.TrackEntryResult_TRACK_ENTRY_RESULT_OK}, nil
}

// GetEntryAuthor returns the author of the entry for passed TrackedEntryKey if the entry is of type Author,
// or the author of the first book in the sequence if the entry type is sequence
func (s *server) GetEntryAuthor(req *proto.TrackedEntryKey) (string, error) {
	log.Printf("GetEntryAuthor: %+v", req)
	if req.EntityId == 0 {
		return "", createRequestError(NoEntryId)
	}

	if req.EntityType == proto.EntryType_ENTRY_TYPE_UNSPECIFIED {
		return "", createRequestError(NoEntryType)
	}

	entryId := req.EntityId

	var author proto.AuthorName
	var err error
	if req.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		author, err = s.db.GetAuthorName(entryId)
		if err != nil {
			return "", err
		}
	} else if req.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
		books, err := s.db.GetSeriesBooks(entryId)
		if err != nil {
			return "", err
		}

		if len(books) == 0 {
			return "", createRequestError(NoEntryId)
		}

		author, err = s.db.GetBookAuthor(int64(books[0].BookId))
		if err != nil {
			return "", err
		}
	} else {
		return "", createRequestError(NoEntryType)
	}

	return FormatAuthorName(author), nil
}

// GetBooksForKey returns all books belonging to author if the passed TrackedEntryKey is of type author,
// or for series, if the key is of type series.
func (s *server) GetBooksForKey(req *proto.TrackedEntryKey) ([]*proto.Book, error) {
	log.Printf("GetBooksForKey: %+v", req)
	if req.EntityId == 0 {
		return nil, createRequestError(NoEntryId)
	}

	if req.EntityType == proto.EntryType_ENTRY_TYPE_UNSPECIFIED {
		return nil, createRequestError(NoEntryType)
	}

	entryId := req.EntityId

	if req.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		return s.db.GetAuthorBooks(entryId)
	} else if req.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
		return s.db.GetSeriesBooks(entryId)
	} else {
		return nil, createRequestError(NoEntryType)
	}
}

// GetEntryName returns the name of the entry for passed TrackedEntryKey
func (s *server) GetEntryName(req *proto.TrackedEntryKey) (string, error) {
	log.Printf("GetEntryName: %+v", req)
	if req.EntityId == 0 {
		return "", createRequestError(NoEntryId)
	}

	if req.EntityType == proto.EntryType_ENTRY_TYPE_UNSPECIFIED {
		return "", createRequestError(NoEntryType)
	}

	entryId := req.EntityId

	if req.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		authorName, err := s.db.GetAuthorName(entryId)
		if err != nil {
			return "", err
		} else {
			return FormatAuthorName(authorName), nil
		}
	} else if req.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
		return s.db.GetSequenceName(entryId)
	} else {
		return "", createRequestError(NoEntryType)
	}
}

func FormatAuthorName(authorName proto.AuthorName) string {
	return fmt.Sprintf("%s, %s %s", authorName.LastName, authorName.FirstName, authorName.MiddleName)
}

func (s *server) ListTrackedEntries(ctx context.Context, req *proto.ListTrackedEntriesRequest) (*proto.ListTrackedEntriesResponse, error) {
	log.Printf("ListTrackedEntries: %+v", req)
	entries := make([]*proto.TrackedEntry, 0)
	err := s.data.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		bprefix := []byte(TrackedEntryPrefix)

		for it.Seek(bprefix); it.ValidForPrefix(bprefix); it.Next() {
			marshalledValue := []byte{}
			key := &proto.KvRecordKey{}

			str := string(it.Item().Key())
			skey := strings.TrimPrefix(str, TrackedEntryPrefix)

			err := proto2.UnmarshalText(skey, key)
			if err != nil {
				return err
			}

			recordKey := key.GetTrackedEntryKey()
			if key == nil {
				continue
			}

			// This isn't really efficient, we should use prefix scan
			if recordKey.UserId != req.UserId {
				continue
			}

			err = it.Item().Value(func(val []byte) error {
				marshalledValue = val
				return nil
			})
			if err != nil {
				return err
			}

			trackedEntry := proto.TrackedEntry{}
			err = proto2.Unmarshal(marshalledValue, &trackedEntry)

			if err != nil {
				return err
			}
			trackedEntry.Key = recordKey
			entries = append(entries, &trackedEntry)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Saved.Seconds*1000+int64(entries[i].Saved.Nanos) >
			entries[j].Saved.Seconds*1000+int64(entries[i].Saved.Nanos)
	})

	return &proto.ListTrackedEntriesResponse{Entry: entries}, nil
}

func (s *server) UntrackEntry(ctx context.Context, req *proto.UntrackEntryRequest) (*proto.UntrackEntryResponse, error) {
	key := req.Key
	log.Printf("UntrackEntry: %+v", key)
	err := s.data.Update(func(txn *badger.Txn) error {
		entryKey := &proto.KvRecordKey{Key: &proto.KvRecordKey_TrackedEntryKey{TrackedEntryKey: key}}
		skey := TrackedEntryPrefix + proto2.MarshalTextString(entryKey)
		key := []byte(skey)

		return txn.Delete(key)
	})
	if err != nil {
		return nil, err
	} else {
		return &proto.UntrackEntryResponse{Key: key, Result: proto.UntrackEntryResult_UNTRACK_ENTRY_RESULT_OK}, nil
	}
}
