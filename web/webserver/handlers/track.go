package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
)

type ActionType int8

const (
	Track ActionType = iota
	Untrack
)

func TrackUntrackHandler(client ClientContext, actionType ActionType) func(ctx *fiber.Ctx) error {
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

		if actionType == Track {
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
		} else if actionType == Untrack {
			resp, err := client.RpcClient.UntrackEntry(&proto.UntrackEntryRequest{
				Key: &proto.TrackedEntryKey{
					UserId:     handlers.MakeUserKeyFromUserNameAndId(userInfo.UserName, userInfo.Id),
					EntityType: proto.EntryType(entityType),
					EntityId:   int64(entityId),
				},
			})

			if err != nil {
				return ctx.Status(500).SendString(err.Error())
			}

			if resp.Result == proto.UntrackEntryResult_UNTRACK_ENTRY_RESULT_OK {
				return ctx.Redirect("/")
			} else if resp.Result == proto.UntrackEntryResult_UNTRACK_ENTRY_RESULT_NOT_TRACKED {
				log.Printf("Entry is not tracked: %d/%d", entityId, entityType)
				return ctx.Redirect("/")
			}
			return ctx.Status(500).SendString(fmt.Sprintf("Unrecognised return code: %d", resp.Result))
		} else {
			return ctx.Status(500).SendString(fmt.Sprintf("Unrecognised action: %d", actionType))
		}
	}
}
