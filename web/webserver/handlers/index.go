package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

func IndexHandler(client ClientContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		ui := c.Locals("user")
		userInfo := ui.(*tgauth.UserInfo)
		sr, err := GetTrackedEntries(client, userInfo)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Render("index", fiber.Map{
			"Name":           userInfo.FirstName,
			"TrackedEntries": sr,
			"BotName":        client.Opts.BotName,
			"DomainName":     client.Opts.DomainName,
		})
	}
}
