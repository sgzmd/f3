package main

import (
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
	first := mockClient.EXPECT().ListTrackedEntries(&proto.ListTrackedEntriesRequest{UserId: "user_123"}).Return(&proto.ListTrackedEntriesResponse{
		Entry: []*proto.TrackedEntry{
			{
				Key: &proto.TrackedEntryKey{
					EntityType: proto.EntryType_ENTRY_TYPE_AUTHOR,
					EntityId:   1,
					UserId:     "user_123",
				},
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
				Key: &proto.TrackedEntryKey{
					EntityType: proto.EntryType_ENTRY_TYPE_AUTHOR,
					EntityId:   1,
					UserId:     "user_123",
				},
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
					Key: &proto.TrackedEntryKey{
						EntityType: proto.EntryType_ENTRY_TYPE_AUTHOR,
						EntityId:   1,
						UserId:     "user_123",
					},
					EntryName:   "Some Entry Name",
					EntryAuthor: "Some Entry Author",
				},
				NewNumEntries: 2,
				NewBook:       []*proto.Book{{BookName: "Some book2 ", BookId: 3234}},
			},
		}}, nil)

	ctx.RpcClient = mockClient

	updates, err := CheckUpdates(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(updates))
	assert.Equal(t, int64(123), updates[0].UserId)
	assert.Equal(t, `<b>Найдены обновления</b>

<b>Some Entry Name</b>

<i>Some book2 </i>
`, updates[0].Message)
}
