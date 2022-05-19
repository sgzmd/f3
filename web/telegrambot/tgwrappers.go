package telegrambot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type IBotApiWrapper interface {
	Send(msg tgbotapi.MessageConfig)
}

type BotApiWrapper struct {
	Bot *tgbotapi.BotAPI
}

func (w BotApiWrapper) Send(msg tgbotapi.MessageConfig) {
	w.Bot.Send(msg)
}
