package updates

import (
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/telegrambot"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"html/template"
	"log"
	"time"
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

const TEMPLATE = `<b>Найдены обновления</b>
{{ range .Updates }}
<b>{{ .EntityName }}</b>
{{ range .Books }}
<i>{{.BookName}}</i>
{{end}}{{ end }}`

func CheckUpdatesLoop(ctx handlers.ClientContext, token string) {
	for {
		err := CheckAndSendUpdates(ctx, token)
		if err != nil {
			log.Printf("Error checking/sending updates: %s", err)
		}
		time.Sleep(time.Minute * 60)
	}
}

func CheckAndSendUpdates(ctx handlers.ClientContext, token string) error {
	updates, err := checkUpdates(ctx)
	if err != nil {
		log.Printf("Error checking updates: %s", err)
		return err
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Error creating bot: %s", err)
		return err
	}
	return sendUpdates(updates, telegrambot.BotApiWrapper{Bot: bot})
}

func checkUpdates(ctx handlers.ClientContext) ([]UpdateMessage, error) {
	resp, err := ctx.RpcClient.ListUsers(&pb.ListUsersRequest{})
	if err != nil {
		return nil, err
	}

	updates := make([]UpdateMessage, 0, 10)

	tmpl, err := template.New("updates").Parse(TEMPLATE)
	if err != nil {
		return nil, err
	}

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

		if len(r2.UpdateRequired) == 0 {
			return nil, nil
		}

		for _, ur := range r2.UpdateRequired {
			upd := Update{
				EntityName: ur.TrackedEntry.EntryName,
				Books:      ur.NewBook,
			}
			update.Updates = append(update.Updates, upd)

			_, e := ctx.RpcClient.TrackEntry(&pb.TrackEntryRequest{
				Key:         ur.TrackedEntry.Key,
				ForceUpdate: true,
			})
			if e != nil {
				log.Printf("Failed to force update %+v: %s", ur.TrackedEntry, e)
			} else {
				log.Printf("Forced update %+v", ur.TrackedEntry)
			}
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

func sendUpdates(updates []UpdateMessage, wrapper telegrambot.IBotApiWrapper) error {
	for _, update := range updates {
		msg := tgbotapi.NewMessage(update.UserId, update.Message)
		msg.ParseMode = "HTML"
		err := wrapper.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
