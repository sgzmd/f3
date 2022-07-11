package handlers

import (
	"fmt"
	"github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"

	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
)

func entryTypeToChar(et proto.EntryType) string {
	if et == proto.EntryType_ENTRY_TYPE_SERIES {
		return "s"
	} else if et == proto.EntryType_ENTRY_TYPE_AUTHOR {
		return "a"
	} else {
		log.Fatalf("Unknown entry type: %v", et)
		return ""
	}
}

type SearchResultEntry struct {
	Entry *proto.FoundEntry
}

type TrackedEntry struct {
	Entry *proto.TrackedEntry
}

func (entry *SearchResultEntry) GetUrl() string {
	return formatFlibustaLink(entry.Entry.EntryType, entry.Entry.EntryId)
}

func (entry *TrackedEntry) GetUrl() string {
	return formatFlibustaLink(entry.Entry.Key.EntityType, entry.Entry.Key.EntityId)
}

func formatFlibustaLink(entryType proto.EntryType, entryId int64) string {
	return fmt.Sprintf("http://flibusta.is/%s/%d",
		entryTypeToChar(entryType), entryId)
}

func (entry *SearchResultEntry) GetEntryClass() string {
	entryType := entry.Entry.EntryType
	return getEntityTypeString(entryType)
}

func (entry *TrackedEntry) GetEntryClass() string {
	return getEntityTypeString(entry.Entry.Key.GetEntityType())
}

func getEntityTypeString(entryType proto.EntryType) string {
	if entryType == proto.EntryType_ENTRY_TYPE_SERIES {
		return "series"
	} else if entryType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		return "author"
	} else {
		return ""
	}
}

func GetTrackedEntries(client ClientContext, userInfo *tgauth.UserInfo) ([]TrackedEntry, error) {
	resp, err := client.RpcClient.ListTrackedEntries(&proto.ListTrackedEntriesRequest{
		UserId: handlers.MakeUserKeyFromUserNameAndId(userInfo.UserName, userInfo.Id),
	})

	if err != nil {
		return nil, err
	}

	sr := make([]TrackedEntry, len(resp.Entry))
	for i, entry := range resp.Entry {
		sr[i] = TrackedEntry{
			Entry: entry,
		}
	}

	return sr, err
}
