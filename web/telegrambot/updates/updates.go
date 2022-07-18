package updates

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sgzmd/f3/web/rpc"
	intrf "github.com/sgzmd/f3/web/telegrambot/intf"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"github.com/sgzmd/f3/web/webserver/updates"
	"log"
)

func CheckUpdatesHandler(update tgbotapi.Update, client rpc.ClientInterface, bot intrf.IBotApiWrapper, token string) {
	n, e := updates.CheckAndSendUpdates(handlers.ClientContext{RpcClient: client}, token)

	if e != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error sending updates: %+v")
		msg.Text = fmt.Sprintf("Error: %+v", e)
		log.Print(msg.Text)
		bot.Send(msg)
		return
	}

	if n == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отсутствие новостей - лучшая новость!")
		log.Print(msg.Text)
		bot.Send(msg)
		return
	}
}
