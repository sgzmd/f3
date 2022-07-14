package main

import (
	"log"

	tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/f3/web/telegrambot"
)

func main() {

	var opts common.Options
	_, err := flags.Parse(&opts)

	if err != nil {
		return
	}

	bot, err := tb.NewBotAPI(opts.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	pclient, err := rpc.NewClient(&opts.GrpcBackend)
	client := pclient
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	botw := &telegrambot.BotApiWrapper{Bot: bot}

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case telegrambot.StartCommand:
					msg := tb.NewMessage(update.Message.Chat.ID, "Hello world!")
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				case telegrambot.SearchCommand:
					telegrambot.SearchCommandHandler(update, client, bot)
				case telegrambot.ListCommand:
					telegrambot.ListCommandHandler(update, client, botw)
				case telegrambot.CheckUpdates:
					telegrambot.CheckUpdatesHandler(update, client, telegrambot.BotApiWrapper{Bot: bot})
				}

			}
		} else if update.CallbackQuery != nil {
			telegrambot.HandleCallbackQuery(update, bot, client)
		}
	}
}
