package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

func IndexHandler(client ClientContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		ui := c.Locals("user")
		userInfo := ui.(*tgauth.UserInfo)
		resp, err := client.RpcClient.ListTrackedEntries(&proto.ListTrackedEntriesRequest{
			UserId: handlers.MakeUserKeyFromUserNameAndId(userInfo.UserName, userInfo.Id),
		})

		log.Printf("Response: %+v", resp)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Render("index", fiber.Map{
			"Name":           userInfo.FirstName,
			"TrackedEntries": resp.Entry,
		})
	}
}
