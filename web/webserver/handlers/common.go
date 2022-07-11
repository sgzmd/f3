package handlers

import (
	"fmt"
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

type ResultEntry struct {
	Entry *proto.FoundEntry
}

func (entry *ResultEntry) GetFlibustaUrl() string {
	return fmt.Sprintf("http://flibusta.is/%s/%d",
		entryTypeToChar(entry.Entry.EntryType), entry.Entry.EntryId)
}

func (entry *ResultEntry) GetEntryClass() string {
	if entry.Entry.EntryType == proto.EntryType_ENTRY_TYPE_SERIES {
		return "series"
	} else if entry.Entry.EntryType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		return "author"
	} else {
		return ""
	}
}
