package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

func ArchiveViewHandler(client ClientContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ui := c.Locals("user")
		userInfo := ui.(*tgauth.UserInfo)
		sr, err := GetTrackedEntries(client, userInfo)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		archivedEntries := make([]TrackedEntry, 0)
		for _, entry := range sr {
			if entry.Entry.Status == proto.TrackedEntryStatus_TRACKED_ENTRY_STATUS_ARCHIVED {
				archivedEntries = append(archivedEntries, entry)
			}
		}

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Render("archive", fiber.Map{
			"Name":           userInfo.FirstName,
			"TrackedEntries": archivedEntries,
			"BotName":        client.Opts.BotName,
			"DomainName":     client.Opts.DomainName,
		})
	}
}
