package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

type SearchResultEntry struct {
	Entry *proto.FoundEntry
}

func (entry *SearchResultEntry) GetFlibustaUrl() string {
	return fmt.Sprintf("http://flibusta.is/%s/%d",
		entryTypeToChar(entry.Entry.EntryType), entry.Entry.EntryId)
}

func (entry *SearchResultEntry) GetEntryClass() string {
	if entry.Entry.EntryType == proto.EntryType_ENTRY_TYPE_SERIES {
		return "series"
	} else if entry.Entry.EntryType == proto.EntryType_ENTRY_TYPE_AUTHOR {
		return "author"
	} else {
		return ""
	}
}

func SearchHandler(client ClientContext) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		searchTerm, _ := url.QueryUnescape(ctx.Params("searchTerm", ""))
		userInfo := ctx.Locals("user").(*tgauth.UserInfo)

		resp, err := client.RpcClient.GlobalSearch(&proto.GlobalSearchRequest{
			SearchTerm: searchTerm,
		})

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		sr := make([]SearchResultEntry, len(resp.Entry))
		for i, entry := range resp.Entry {
			sr[i] = SearchResultEntry{
				Entry: entry,
			}
		}

		return ctx.Render("index", fiber.Map{
			"Name": userInfo.FirstName, "HasSearchResults": true, "SearchResults": sr,
		})
	}
}
