package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/mocks"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestCheckUpdatesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mocks.NewMockClientInterface(ctrl)
	bot := mocks.NewMockIBotApiWrapper(ctrl)

	update := NewFakeUpdate()

	listResp := NewFakeListTrackedEntriesResponse(5, update.Message.From.UserName)
	client.
		EXPECT().
		ListTrackedEntries(gomock.Any()).
		Return(listResp, nil)

	client.
		EXPECT().
		CheckUpdates(&pb.CheckUpdatesRequest{TrackedEntry: listResp.Entry}).
		Return(NewFakeCheckUpdatesResponse(), nil)

	client.EXPECT().
		TrackEntry(&pb.TrackEntryRequest{
			Key: &pb.TrackedEntryKey{
				EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
				EntityId:   1,
				UserId:     "testuser",
			},
			ForceUpdate: true,
		})

	bot.EXPECT().Send(gomock.Any())

	CheckUpdatesHandler(update, client, bot)
}

func NewFakeCheckUpdatesResponse() *pb.CheckUpdatesResponse {
	entries := make([]*pb.UpdateRequired, 1)
	entries[0] = &pb.UpdateRequired{
		TrackedEntry: &pb.TrackedEntry{
			Key: &pb.TrackedEntryKey{
				EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
				EntityId:   1,
				UserId:     "testuser",
			},
			EntryName:   "Test Entry Name",
			NumEntries:  12,
			EntryAuthor: "",
			Book:        nil,
			Saved: &timestamppb.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
		},
		NewNumEntries: 0,
		NewBook:       nil,
	}

	return &pb.CheckUpdatesResponse{
		UpdateRequired: entries,
	}
}

// creates deep fake pb.ListTrackedEntriesResponse
func NewFakeListTrackedEntriesResponse(n int, userId string) *pb.ListTrackedEntriesResponse {
	entries := make([]*pb.TrackedEntry, n)
	for i := 0; i < n; i++ {
		entries[i] = NewFakeTrackedEntry(userId, i, pb.EntryType_ENTRY_TYPE_SERIES)
	}
	return &pb.ListTrackedEntriesResponse{
		Entry: entries,
	}
}

// creates deep fake pb.TrackedEntry
func NewFakeTrackedEntry(userId string, entryId int, entryType pb.EntryType) *pb.TrackedEntry {
	return &pb.TrackedEntry{
		Key: &pb.TrackedEntryKey{
			EntityType: entryType,
			EntityId:   int64(entryId),
			UserId:     userId,
		},
		EntryName:   "",
		NumEntries:  12,
		EntryAuthor: "",
		Book:        nil,
		Saved: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
	}
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
			Chat:                 &tgbotapi.Chat{ID: 123456},
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
