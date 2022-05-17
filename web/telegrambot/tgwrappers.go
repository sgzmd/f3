package telegrambot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type IBotApiWrapper interface {
	Send(msg tgbotapi.MessageConfig)
}
