package handlers

import (
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEntryUrl(t *testing.T) {
	assert.Equal(t, "http://flibusta.is/s/123", TrackedEntryUrl(&proto.TrackedEntry{
		Key: &proto.TrackedEntryKey{
			EntityType: proto.EntryType_ENTRY_TYPE_SERIES,
			EntityId:   123,
			UserId:     "test",
		},
		EntryName:   "",
		NumEntries:  0,
		EntryAuthor: "",
		Book:        nil,
		Saved:       nil,
	}))

	assert.Equal(t, "http://flibusta.is/a/123", TrackedEntryUrl(&proto.TrackedEntry{
		Key: &proto.TrackedEntryKey{
			EntityType: proto.EntryType_ENTRY_TYPE_AUTHOR,
			EntityId:   123,
			UserId:     "test",
		},
		EntryName:   "",
		NumEntries:  0,
		EntryAuthor: "",
		Book:        nil,
		Saved:       nil,
	}))
}
