package main

import (
	"context"

	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"

	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

var (
	TrackableEntries = map[int64]pb.EntryType{
		1801:   pb.EntryType_ENTRY_TYPE_AUTHOR,
		109170: pb.EntryType_ENTRY_TYPE_AUTHOR,
		34145:  pb.EntryType_ENTRY_TYPE_SERIES,
	}
)

type FlibustierStorageSuite struct {
	suite.Suite
	client pb.FlibustierServiceClient
	conn   *grpc.ClientConn
}

func TestFlibustierStorage(t *testing.T) {
	suite.Run(t, new(FlibustierStorageSuite))
}

func (suite *FlibustierStorageSuite) TestServer_TrackEntry() {
	trackReq := &pb.TrackEntryRequest{
		EntryType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntryId:   109170,
		UserId:    "1"}
	resp, err := suite.client.TrackEntry(context.Background(), trackReq)
	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK, resp.Result)

	// Second time should fail
	resp2, err := suite.client.TrackEntry(context.Background(), trackReq)
	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_ALREADY_TRACKED, resp2.Result)

	resp3, err := suite.client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "1"})
	suite.Assert().Nil(err)
	suite.Assert().Len(resp3.Entry[0].Book, 8)
}

func (suite *FlibustierStorageSuite) TestServer_ListTrackedEntries() {
	entriesToTrack := make(map[int64]pb.EntryType)
	for k, v := range TrackableEntries {
		entriesToTrack[k] = v
	}

	for key, value := range entriesToTrack {
		suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
			EntryId:   key,
			EntryType: value,
			UserId:    "1",
		})
	}

	_, _ = suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
		EntryId:   1801,
		EntryType: entriesToTrack[1801],
		UserId:    "another_user",
	})

	resp, err := suite.client.ListTrackedEntries(
		context.Background(),
		&pb.ListTrackedEntriesRequest{UserId: "1"})
	suite.Assert().Nil(err)

	for _, entry := range resp.Entry {
		suite.Assert().Contains(entriesToTrack, entry.EntryId)
		delete(entriesToTrack, entry.EntryId)
	}

	suite.Assert().Empty(entriesToTrack)
}

func createTrackedEntry(i int, uid string) *pb.TrackEntryRequest {
	return &pb.TrackEntryRequest{
		EntryType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntryId:   int64(i),
		UserId:    uid}
}

func (suite *FlibustierStorageSuite) TestServer_Untrack() {
	var entryId int64
	var entryType pb.EntryType
	for k, v := range TrackableEntries {
		entryId = k
		entryType = v
		break
	}

	r, err := suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
		EntryId:   entryId,
		EntryType: entryType,
		UserId:    "user",
	})
	suite.Assert().Nil(err, err)

	r2, err := suite.client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "user"})
	suite.Assert().Len(r2.Entry, 1)

	suite.client.UntrackEntry(context.Background(), &pb.UntrackEntryRequest{Key: r.Key})

	r3, _ := suite.client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "user"})
	suite.Assert().Empty(r3.Entry)
}

func (suite *FlibustierStorageSuite) BeforeTest(suiteName, testName string) {
	ctx := context.Background()
	// Creating a client
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(FLIBUSTA_DB)))
	if err != nil {
		panic(err)
	}
	suite.client = pb.NewFlibustierServiceClient(conn)
	suite.conn = conn
}

func (suite *FlibustierStorageSuite) AfterTest(suiteName, testName string) {
	suite.conn.Close()
}
