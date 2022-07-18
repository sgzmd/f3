package updates

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc/mocks"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckUpdates(t *testing.T) {
	ctrl := gomock.NewController(t)

	ctx := handlers.ClientContext{}
	mockClient := mocks.NewMockClientInterface(ctrl)
	mockClient.EXPECT().ListUsers(gomock.Any()).Return(&proto.ListUsersResponse{
		User: []*proto.UserInfo{
			{
				UserId:         "user_123",
				UserName:       "test",
				UserTelegramId: 123,
				UserEmail:      "",
			},
			{
				UserId:         "user_456",
				UserName:       "test2",
				UserTelegramId: 456,
				UserEmail:      "",
			},
		},
	}, nil)
	key1 := &proto.TrackedEntryKey{
		EntityType: proto.EntryType_ENTRY_TYPE_AUTHOR,
		EntityId:   1,
		UserId:     "user_123",
	}
	first := mockClient.EXPECT().ListTrackedEntries(&proto.ListTrackedEntriesRequest{UserId: "user_123"}).Return(&proto.ListTrackedEntriesResponse{
		Entry: []*proto.TrackedEntry{
			{
				Key:         key1,
				EntryName:   "Author 1",
				NumEntries:  1,
				EntryAuthor: "Author 1",
				Book: []*proto.Book{{
					BookName: "Some book",
					BookId:   1234,
				}},
				Saved: nil,
			},
		},
	}, nil)

	second := mockClient.EXPECT().ListTrackedEntries(&proto.ListTrackedEntriesRequest{UserId: "user_456"}).Return(&proto.ListTrackedEntriesResponse{
		Entry: []*proto.TrackedEntry{},
	}, nil)

	gomock.InOrder(first, second)
	mockClient.EXPECT().CheckUpdates(&proto.CheckUpdatesRequest{
		TrackedEntry: []*proto.TrackedEntry{
			{
				Key:         key1,
				EntryName:   "Author 1",
				NumEntries:  1,
				EntryAuthor: "Author 1",
				Book: []*proto.Book{{
					BookName: "Some book",
					BookId:   1234,
				}},
				Saved: nil,
			},
		},
	}).Return(&proto.CheckUpdatesResponse{
		UpdateRequired: []*proto.UpdateRequired{
			{
				TrackedEntry: &proto.TrackedEntry{
					Key:         key1,
					EntryName:   "Some Entry Name",
					EntryAuthor: "Some Entry Author",
				},
				NewNumEntries: 2,
				NewBook:       []*proto.Book{{BookName: "Some book2 ", BookId: 3234}},
			},
		}}, nil)

	mockClient.EXPECT().TrackEntry(&proto.TrackEntryRequest{
		Key:         key1,
		ForceUpdate: true,
	})

	ctx.RpcClient = mockClient

	updates, err := checkUpdates(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(updates))
	assert.Equal(t, int64(123), updates[0].UserId)
	assert.Equal(t, `<b>Найдены обновления</b>

<b>Some Entry Name</b>

<i>Some book2 </i>
`, updates[0].Message)
}

func TestSendUpdates(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClient := mocks.NewMockIBotApiWrapper(ctrl)

	msg := UpdateMessage{
		UserId:  123,
		Message: "Some message",
	}
	tgmessage := tgbotapi.NewMessage(msg.UserId, msg.Message)
	tgmessage.ParseMode = "HTML"
	mockClient.EXPECT().Send(tgmessage).Return(nil)

	_, err := sendUpdates([]UpdateMessage{msg}, mockClient)

	assert.Nil(t, err)
}
