package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
)

func TrackHandler(client ClientContext) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		entityType, err := ctx.ParamsInt("entityType", -1)
		if err != nil {
			return fiber.ErrBadRequest
		}

		entityId, err := ctx.ParamsInt("id", -1)
		if err != nil {
			return fiber.ErrBadRequest
		}

		ui := ctx.Locals("user")
		userInfo := ui.(*tgauth.UserInfo)

		resp, err := client.RpcClient.TrackEntry(&proto.TrackEntryRequest{
			Key: &proto.TrackedEntryKey{
				UserId:     handlers.MakeUserKeyFromUserNameAndId(userInfo.UserName, userInfo.Id),
				EntityType: proto.EntryType(entityType),
				EntityId:   int64(entityId),
			},
		})

		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		if resp.Result == proto.TrackEntryResult_TRACK_ENTRY_RESULT_OK {
			return ctx.Redirect("/")
		} else {
			// Need to do something else here?
			return ctx.Redirect("/")
		}
	}
}
