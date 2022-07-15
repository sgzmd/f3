package telegrambot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sgzmd/f3/web/rpc"
	"log"
)

func BotFunc(telegramToken string, client rpc.ClientInterface) {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	botw := &BotApiWrapper{Bot: bot}

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case StartCommand:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world!")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				case SearchCommand:
					SearchCommandHandler(update, client, bot)
				case ListCommand:
					ListCommandHandler(update, client, botw)
				case CheckUpdates:
					CheckUpdatesHandler(update, client, BotApiWrapper{Bot: bot})
				}

			}
		} else if update.CallbackQuery != nil {
			HandleCallbackQuery(update, bot, client)
		}
	}
}
