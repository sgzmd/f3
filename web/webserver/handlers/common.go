package handlers

import (
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
