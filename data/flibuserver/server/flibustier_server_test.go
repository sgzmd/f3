package main

import (
	"context"
	"flag"
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
	actionCreate := pb.UserInfoAction_USER_INFO_ACTION_CREATE
	req := &pb.GetUserInfoRequest{
		UserId:         "123",
		UserTelegramId: 1234,
		Action:         &actionCreate,
	}
	resp, err := client.GetUserInfo(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, "123", resp.UserInfo.UserId)
	assert.Equal(t, int64(1234), resp.UserInfo.UserTelegramId)
}

// Tests that user can be created, and then retrieved
func TestServer_TestGetUserInfo_Get(t *testing.T) {
	actionCreate := pb.UserInfoAction_USER_INFO_ACTION_CREATE
	// Create the user first
	req := &pb.GetUserInfoRequest{
		UserId:         "567",
		UserTelegramId: 78931,
		Action:         &actionCreate,
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

// Tests ListUsers RPC
func TestServer_TestListUsers(t *testing.T) {
	_, e := client.DeleteAllUsers(context.Background(), &pb.DeleteAllUsersRequest{})
	assert.Nil(t, e)

	actionCreate := pb.UserInfoAction_USER_INFO_ACTION_CREATE
	req := &pb.GetUserInfoRequest{
		UserId:         "567",
		UserTelegramId: 78931,
		Action:         &actionCreate,
	}
	resp, err := client.GetUserInfo(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, "567", resp.UserInfo.UserId)

	r, e := client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	assert.Nil(t, e)
	assert.Len(t, r.User, 1)

	req2 := &pb.GetUserInfoRequest{
		UserId:         "981",
		UserTelegramId: 789321,
		Action:         &actionCreate,
	}
	_, _ = client.GetUserInfo(context.Background(), req2)

	r, e = client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	assert.Nil(t, e)
	assert.Len(t, r.User, 2)

	_, _ = client.DeleteAllUsers(context.Background(), &pb.DeleteAllUsersRequest{})

	r, e = client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	assert.Nil(t, e)
	assert.Len(t, r.User, 0)
}

func TestMain(m *testing.M) {
	flag.Parse()

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
