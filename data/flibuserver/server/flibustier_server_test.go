package main

import (
	"context"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const (
	FLIBUSTA_DB = "../../../testutils/flibusta-test.db"
)

var (
	client pb.FlibustierServiceClient = nil
)

func TestSmokeTest(t *testing.T) {
	const TERM = "метел"
	result, err := client.GlobalSearch(context.Background(), &pb.GlobalSearchRequest{SearchTerm: TERM})
	if err != nil {
		log.Fatalf("Failed the smoke test: %v", err)
		t.Fatalf("Smoke test failed")
	} else {
		assert.Equal(t, TERM, result.OriginalRequest.SearchTerm, "Request doesn't look right")
	}
}

func TestSearchEverything(t *testing.T) {
	const TERM = "Николай Александрович Метельский"
	result, err := client.GlobalSearch(context.Background(), &pb.GlobalSearchRequest{SearchTerm: TERM})
	assert.Nil(t, err)
	assert.Len(t, result.Entry, 2)
	assert.Equal(t, result.Entry[0].Author, "Николай Александрович Метельский")
	assert.Equal(t, result.Entry[1].EntryName, "Унесенный ветром")
}

func TestSearchAuthor(t *testing.T) {
	const TERM = "Метельский"
	result, err := client.GlobalSearch(context.Background(),
		&pb.GlobalSearchRequest{SearchTerm: TERM, EntryTypeFilter: pb.EntryType_ENTRY_TYPE_AUTHOR})
	assert.Nil(t, err)
	assert.Len(t, result.Entry, 1)
	assert.Equal(t, result.Entry[0].Author, "Николай Александрович Метельский")
}

func TestSearchSeries(t *testing.T) {
	// note that searching for author of the series
	const TERM = "Метельский"
	result, err := client.GlobalSearch(context.Background(),
		&pb.GlobalSearchRequest{SearchTerm: TERM, EntryTypeFilter: pb.EntryType_ENTRY_TYPE_SERIES})
	assert.Nil(t, err)
	assert.Len(t, result.Entry, 1)
	assert.Equal(t, result.Entry[0].EntryName, "Унесенный ветром")
}

func TestCheckUpdates_Author(t *testing.T) {
	books := []*pb.Book{{BookId: 452501, BookName: "Чужие маски"}}

	tracked := &pb.TrackedEntry{Key: &pb.TrackedEntryKey{
		EntityType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntityId:   109170,
		UserId:     "123",
	}, EntryName: "Метельский", NumEntries: 1, Book: books}

	request := pb.CheckUpdatesRequest{
		TrackedEntry: []*pb.TrackedEntry{tracked},
	}

	resp, err := client.CheckUpdates(context.Background(), &request)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	} else {
		// t.Errorf("Result: %s", resp.String())
		if len(resp.UpdateRequired) != 1 || resp.UpdateRequired[0].NewNumEntries != 9 {
			t.Fatalf(
				"Expect to have 1 UpdateRequired entity with 9 new_num_entries, but have: %s",
				resp)
		}
	}
}

func TestCheckUpdates_Series(t *testing.T) {
	books := []*pb.Book{{BookId: 452501, BookName: "Чужие маски"}}

	tracked := &pb.TrackedEntry{Key: &pb.TrackedEntryKey{
		EntityType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntityId:   109170,
		UserId:     "123",
	}, EntryName: "Метельский", NumEntries: 1, Book: books}

	request := pb.CheckUpdatesRequest{
		TrackedEntry: []*pb.TrackedEntry{tracked},
	}

	resp, err := client.CheckUpdates(context.Background(), &request)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	} else {
		// t.Errorf("Result: %s", resp.String())
		if len(resp.UpdateRequired) != 1 || resp.UpdateRequired[0].NewNumEntries != 9 {
			t.Fatalf(
				"Expect to have 1 UpdateRequired entity with 9 new_num_entries, but have: %s",
				resp)
		}
	}
}

func TestServer_GetSeriesBooks(t *testing.T) {
	req := &pb.GetSeriesBooksRequest{SequenceId: 34145}
	resp, err := client.GetSeriesBooks(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed: %+v", err)
	} else {
		assert.Equal(t, req.SequenceId, resp.EntityBookResponse.EntityId)
		assert.Len(t, resp.EntityBookResponse.Book, 8)

		assert.Equal(t, "Унесенный ветром: Меняя маски. Теряя маски. Чужие маски",
			resp.EntityBookResponse.Book[0].BookName)
		assert.Equal(t, "Унесенный ветром", resp.EntityBookResponse.EntityName.GetSequenceName())
	}
}

func TestServer_GetAuthorBooks(t *testing.T) {
	req := &pb.GetAuthorBooksRequest{AuthorId: 109170}
	resp, err := client.GetAuthorBooks(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed: %+v", err)
	} else {
		assert.Equal(t, req.AuthorId, resp.EntityBookResponse.EntityId)
		assert.Len(t, resp.EntityBookResponse.Book, 8)

		assert.Equal(t, "Чужие маски", resp.EntityBookResponse.Book[0].BookName)
		assert.Equal(t, &pb.AuthorName{
			LastName:   "Метельский",
			MiddleName: "Александрович",
			FirstName:  "Николай"}, resp.EntityBookResponse.EntityName.GetAuthorName())
	}
}

func TestServer_TestGetUserInfo_NotFound(t *testing.T) {
	req := &pb.GetUserInfoRequest{UserId: "123"}
	resp, err := client.GetUserInfo(context.Background(), req)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "User not found")
}

func TestServer_TestGetUserInfo_Create(t *testing.T) {
	req := &pb.GetUserInfoRequest{
		UserId:         "123",
		UserTelegramId: 1234,
		Action:         pb.UserInfoAction_USER_INFO_ACTION_CREATE,
	}
	resp, err := client.GetUserInfo(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, "123", resp.UserInfo.UserId)
	assert.Equal(t, int64(1234), resp.UserInfo.UserTelegramId)
}

// Tests that user can be created, and then retrieved
func TestServer_TestGetUserInfo_Get(t *testing.T) {
	// Create the user first
	req := &pb.GetUserInfoRequest{
		UserId:         "567",
		UserTelegramId: 78931,
		Action:         pb.UserInfoAction_USER_INFO_ACTION_CREATE,
	}
	resp, err := client.GetUserInfo(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, "567", resp.UserInfo.UserId)

	// Get the user
	req = &pb.GetUserInfoRequest{
		UserId: "567",
	}
	resp, err = client.GetUserInfo(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, "567", resp.UserInfo.UserId)
	assert.Equal(t, int64(78931), resp.UserInfo.UserTelegramId)
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	// Creating a client
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(FLIBUSTA_DB)))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client = pb.NewFlibustierServiceClient(conn)

	ret := m.Run()

	conn.Close()
	os.Exit(ret)
}
