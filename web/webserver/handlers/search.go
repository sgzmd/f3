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

		resp, err := client.RpcClient.GlobalSearch(&proto.GlobalSearchRequest{
			SearchTerm: searchTerm,
		})

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		sr := make([]ResultEntry, len(resp.Entry))
		for i, entry := range resp.Entry {
			sr[i] = ResultEntry{
				Entry: entry,
			}
		}

		return ctx.Render("index", fiber.Map{
			"Name": userInfo.FirstName, "HasSearchResults": true, "SearchResults": sr,
		})
	}
}
