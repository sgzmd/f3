package main

import (
	"bytes"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"html/template"
	"log"
)

type Update struct {
	EntityName string

	Books []*pb.Book
}

type UpdateMsg struct {
	Updates []Update
}

type UpdateMessage struct {
	UserId  int64
	Message string
}

func CheckUpdates(ctx handlers.ClientContext) ([]UpdateMessage, error) {
	resp, err := ctx.RpcClient.ListUsers(&pb.ListUsersRequest{})
	if err != nil {
		return nil, err
	}

	updates := make([]UpdateMessage, 0, 10)

	tmpl := template.Must(template.ParseFiles("templates/messages/update.html"))

	for _, user := range resp.User {
		update := UpdateMsg{Updates: []Update{}}

		log.Printf("Checking updates for %+v", user)

		resp, err := ctx.RpcClient.ListTrackedEntries(&pb.ListTrackedEntriesRequest{UserId: user.UserId})
		if err != nil {
			return nil, err
		}

		if len(resp.Entry) == 0 {
			continue
		}

		r2, err := ctx.RpcClient.CheckUpdates(&pb.CheckUpdatesRequest{TrackedEntry: resp.Entry})
		if err != nil {
			return nil, err
		}

		for _, ur := range r2.UpdateRequired {
			upd := Update{
				EntityName: ur.TrackedEntry.EntryName,
				Books:      ur.NewBook,
			}
			update.Updates = append(update.Updates, upd)
		}

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, update)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return nil, err
		}

		updates = append(updates, UpdateMessage{UserId: user.UserTelegramId, Message: tpl.String()})
	}

	return updates, nil
}
