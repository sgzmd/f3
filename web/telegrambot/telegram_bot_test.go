package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/sgzmd/f3/web/mocks"
	"testing"
)

func TestCheckUpdatesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mocks.NewMockClientInterface(ctrl)
	bot := mocks.NewMockIBotApiWrapper(ctrl)

	update := NewFakeUpdate()

	CheckUpdatesHandler(update, client, bot)
}

// creates deep fake tgbotapi.Update object
func NewFakeUpdate() tgbotapi.Update {
	return tgbotapi.Update{
		UpdateID: 0,
		Message: &tgbotapi.Message{
			MessageID: 123,
			From: &tgbotapi.User{
				ID:                      123,
				IsBot:                   false,
				FirstName:               "",
				LastName:                "",
				UserName:                "testuser",
				LanguageCode:            "",
				CanJoinGroups:           false,
				CanReadAllGroupMessages: false,
				SupportsInlineQueries:   false,
			},
			SenderChat:           nil,
			Date:                 0,
			Chat:                 nil,
			ForwardFrom:          nil,
			ForwardFromChat:      nil,
			ForwardFromMessageID: 0,
			ForwardSignature:     "",
			ForwardSenderName:    "",
			ForwardDate:          0,
			IsAutomaticForward:   false,
			ReplyToMessage:       nil,
			ViaBot:               nil,
			EditDate:             0,
			HasProtectedContent:  false,
			MediaGroupID:         "",
			AuthorSignature:      "",
			Text:                 "",
			Entities:             nil,
			Animation:            nil,
			Audio:                nil,
			Document:             nil,
			Photo:                nil,
			Sticker:              nil,
			Video:                nil,
			VideoNote:            nil,
			Voice:                nil,
			Caption:              "",
			CaptionEntities:      nil,
			Contact:              nil,
			Dice:                 nil,
			Game:                 nil,
		},
	}
}
