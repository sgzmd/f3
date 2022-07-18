package intf

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotApiWrapper struct {
	Bot *tgbotapi.BotAPI
}

func (w BotApiWrapper) Send(msg tgbotapi.MessageConfig) error {
	_, err := w.Bot.Send(msg)
	return err
}
