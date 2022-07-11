package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

func SearchHandler(client ClientContext) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		searchTerm, _ := url.QueryUnescape(ctx.Params("searchTerm", ""))
		userInfo := ctx.Locals("user").(*tgauth.UserInfo)

		var tr []TrackedEntry
		var err error
		finished := make(chan bool)
		// Getting tracked entries async
		go func() {
			tr, err = GetTrackedEntries(client, userInfo)
			finished <- true
		}()

		resp, err := client.RpcClient.GlobalSearch(&proto.GlobalSearchRequest{
			SearchTerm: searchTerm,
		})

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		<-finished
		sr := make([]SearchResultEntry, len(resp.Entry))
		for i, entry := range resp.Entry {
			sr[i] = SearchResultEntry{
				Entry:   entry,
				Tracked: false,
			}

			for _, trackedEntry := range tr {
				if trackedEntry.Entry.Key.EntityType == entry.EntryType &&
					trackedEntry.Entry.Key.EntityId == entry.EntryId {
					sr[i].Tracked = true
				}
			}
		}

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		return ctx.Render("index", fiber.Map{
			"Name":             userInfo.FirstName,
			"HasSearchResults": true,
			"SearchResults":    sr,
			"TrackedEntries":   tr,
		})
	}
}
