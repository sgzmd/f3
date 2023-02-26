package main

import (
	"context"
	"github.com/dgraph-io/badger/v3"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb/sqlite3"
	"google.golang.org/grpc/test/bufconn"
	"io/ioutil"
	"log"
	"math"
	"net"
	"path/filepath"
	"runtime"
	"time"

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
	trackReq := &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
		EntityType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntityId:   109170,
		UserId:     "1"},
		ForceUpdate: false}
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
	type idType struct {
		id int64
		t  pb.EntryType
	}

	entries := make([]idType, 0)
	entriesToTrack := make(map[int64]pb.EntryType)
	for k, v := range TrackableEntries {
		entries = append(entries, idType{id: k, t: v})
	}

	for i, _ := range entries {
		// Submitting in reverse order to make assertion easier
		val := entries[len(entries)-i-1]
		_, e := suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
			EntityType: val.t,
			EntityId:   val.id,
			UserId:     "1",
		}})
		suite.Assert().Nil(e)
		d, _ := time.ParseDuration("1s")
		time.Sleep(d)
	}

	_, _ = suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
		EntityId:   1801,
		EntityType: entriesToTrack[1801],
		UserId:     "another_user",
	}})

	resp, err := suite.client.ListTrackedEntries(
		context.Background(),
		&pb.ListTrackedEntriesRequest{UserId: "1"})
	suite.Assert().Nil(err)

	// Checking that entries were returned in the order expected
	result := make([]idType, 0)
	for _, entry := range resp.Entry {
		result = append(result, idType{id: entry.Key.EntityId, t: entry.Key.EntityType})
	}

	suite.Assert().Equal(entries, result)
}

func (suite *FlibustierStorageSuite) TestServer_Untrack() {
	var entryId int64
	var entryType pb.EntryType
	for k, v := range TrackableEntries {
		entryId = k
		entryType = v
		break
	}

	r, err := suite.client.TrackEntry(context.Background(), &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
		EntityId:   entryId,
		EntityType: entryType,
		UserId:     "user",
	}})
	suite.Assert().Nil(err, err)

	r2, err := suite.client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "user"})
	suite.Assert().Len(r2.Entry, 1)

	suite.client.UntrackEntry(context.Background(), &pb.UntrackEntryRequest{Key: r.Key})

	r3, _ := suite.client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "user"})
	suite.Assert().Empty(r3.Entry)
}

func (suite *FlibustierStorageSuite) TestServer_TrackEntry_ListTracked() {
	r, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "abc"})
	suite.Assert().Nil(err)
	suite.Assert().NotNil(r)

	req := &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
		EntityId:   109170,
		EntityType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		UserId:     "abc",
	}}
	resp, err := client.TrackEntry(context.Background(), req)

	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK, resp.Result)

	r2, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "abc"})
	suite.Assert().Nil(err)

	suite.Assert().Len(r2.Entry, 1)
	suite.Assert().Equal(int64(109170), r2.Entry[0].Key.EntityId)
	suite.Assert().LessOrEqual(math.Abs(float64(r2.Entry[0].Saved.Seconds)-float64(time.Now().Unix())), float64(2))
	suite.Assert().Equal("Метельский, Николай Александрович", r2.Entry[0].EntryAuthor)
}

func (suite *FlibustierStorageSuite) TestServer_TrackEntry_Double() {
	theKey := &pb.TrackedEntryKey{
		EntityId:   34145,
		EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
		UserId:     "abc"}
	req := &pb.TrackEntryRequest{Key: theKey, ForceUpdate: false}
	resp, err := client.TrackEntry(context.Background(), req)
	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK, resp.Result)

	// Still cannot figure how to recreate the state for each test run
	defer client.UntrackEntry(context.Background(), &pb.UntrackEntryRequest{Key: theKey})

	resp, err = client.TrackEntry(context.Background(), req)
	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_ALREADY_TRACKED, resp.Result)

	req2 := &pb.TrackEntryRequest{Key: theKey, ForceUpdate: true}
	resp, err = client.TrackEntry(context.Background(), req2)
	suite.Assert().Nil(err)
	suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK, resp.Result)
}

func (suite *FlibustierStorageSuite) TestCheckUpdates_Author() {
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
		suite.Failf("CheckUpdates failed: %+v", err.Error())
	} else {
		suite.Assert().Len(resp.UpdateRequired, 1)
		suite.Assert().Equal(int32(8), resp.UpdateRequired[0].NewNumEntries)
	}
}

func (suite *FlibustierStorageSuite) TestCheckUpdates_Series() {
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
		suite.Failf("Failed: %v", err.Error())
	} else {
		// t.Errorf("Result: %s", resp.String())
		suite.Assert().Len(resp.UpdateRequired, 1)
		suite.Assert().Equal(int32(8), resp.UpdateRequired[0].NewNumEntries)
	}
}

func (suite *FlibustierStorageSuite) TestCheckUpdates_Retrack() {
	client.DeleteAllTracked(context.Background(), &pb.DeleteAllTrackedRequest{})

	books := []*pb.Book{{BookId: 452501, BookName: "Чужие маски"}}

	tracked := &pb.TrackedEntry{Key: &pb.TrackedEntryKey{
		EntityType: pb.EntryType_ENTRY_TYPE_AUTHOR,
		EntityId:   109170,
		UserId:     "123",
	}, EntryName: "Метельский", NumEntries: 1, Book: books}

	checkUpdatesRequest := pb.CheckUpdatesRequest{
		TrackedEntry: []*pb.TrackedEntry{tracked},
	}

	resp, err := client.CheckUpdates(context.Background(), &checkUpdatesRequest)
	if err != nil {
		suite.Fail("Failed: %+v", err)
	} else {
		// t.Errorf("Result: %s", resp.String())
		suite.Assert().Len(resp.UpdateRequired, 1)
		suite.Assert().Equal(int32(8), resp.UpdateRequired[0].NewNumEntries)

		// Retrack the same entry so the new books are added
		resp, err := client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
			Key:         tracked.Key,
			ForceUpdate: true,
		})

		if err != nil {
			suite.Fail("Failed force-track: %+v", err)
		}

		suite.Assert().Equal(pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK, resp.Result)

		// Listing entry agan
		listResp, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{
			UserId: "123",
		})
		if err != nil {
			suite.Fail("Failed listing entries: %+v", err)
		}

		var newTrackedEntry *pb.TrackedEntry
		for idx, tracked := range listResp.Entry {
			if tracked.Key.EntityId == 109170 {
				newTrackedEntry = listResp.Entry[idx]
				break
			}
		}
		if newTrackedEntry == nil {
			suite.Fail("Failed to find tracked entry")
		}

		checkUpdatesRequest.TrackedEntry[0] = newTrackedEntry
		r2, err := client.CheckUpdates(context.Background(), &checkUpdatesRequest)
		if err != nil {
			suite.Fail("Failed: %+v", err)
		} else {
			if len(r2.UpdateRequired) > 0 {
				suite.Failf("Expect to have no UpdateRequired entity, but have: %s", r2.String())
			}
		}
	}
}

func (suite *FlibustierStorageSuite) BeforeTest(suiteName, testName string) {
	log.Print("BeforeTest()")
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
	log.Print("AfterTest()")
	suite.conn.Close()
}

func newServerWithDump(db_path string, datastore string, dump string) (*server, error) {
	srv := new(server)

	db, err := OpenDatabase(db_path)
	if err != nil {
		return nil, err
	}
	db.Exec(dump)

	srv.db = sqlite3.NewSqlite3Db(db)

	var opt badger.Options
	if datastore == "" {
		opt = badger.DefaultOptions("").WithInMemory(true)
	} else {
		opt = badger.DefaultOptions(datastore)
	}

	srv.data, err = badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func dialer(flibustaDb string) func(context.Context, string) (net.Conn, error) {
	log.Print("dialer() creates new server")
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(filename)

	data, err := ioutil.ReadFile(dir + "../../../testutils/flibusta-test-db.sql")
	if err != nil {
		panic(err)
	}
	sqlDump := string(data)
	srv, err := newServerWithDump(flibustaDb, "" /* in memory datastore */, sqlDump)
	if err != nil {
		panic(err)
	}
	pb.RegisterFlibustierServiceServer(server, srv)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
