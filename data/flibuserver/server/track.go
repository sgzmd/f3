package main

import (
	"context"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	proto2 "github.com/golang/protobuf/proto"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb/sqlite3"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sort"
	"strings"
	"time"
)

func (s *server) TrackEntry(ctx context.Context, req *proto.TrackEntryRequest) (*proto.TrackEntryResponse, error) {
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
	var entries []*proto.FoundEntry
	var bookExtractorSql string
	entryId := int(key.EntityId)

	if key.EntityType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		query := sqlite3.CreateAuthorByIdQuery(entryId)
		log.Printf("SQL: %s", query)
		entries, err = s.iterateOverAuthors(query)
		bookExtractorSql = sqlite3.CreateGetBooksForAuthor(entryId)
	} else if key.EntityType == proto.EntryType_ENTRY_TYPE_SERIES {
		query := sqlite3.CreateSequenceByIdQuery(entryId)
		log.Printf("SQL: %s", query)
		entries, err = s.iterateOverSeries(query)
		bookExtractorSql = sqlite3.CreateGetBooksBySequenceId(entryId)
	} else {
		return nil, createRequestError(NoEntryType)
	}

	if err != nil {
		log.Printf("Error requesting entries for entity: %+v", err)
		return nil, err
	}

	if len(entries) != 1 {
		e := fmt.Errorf("must be 1 entry exactly in %s", entries)
		log.Printf("Error: %+v", e)
		return nil, e
	}

	entry := entries[0]
	books, err := s.getBooks(bookExtractorSql)
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

		val := proto.TrackedEntry{
			Key:         key,
			EntryName:   entry.EntryName,
			NumEntries:  entry.NumEntities,
			Book:        books,
			EntryAuthor: entry.Author,
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
