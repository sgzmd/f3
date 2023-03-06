package integration

import (
	"context"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
)

// This test runs in assumption that GRPC backend was already brought up with all of the dependencies

var client proto.FlibustierServiceClient

func TestSmokeTest(t *testing.T) {
	resp, err := client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	if err != nil {
		t.Fatalf("ListUsers failed: %v", err)
	}
	t.Logf("ListUsers response: %+v", resp)
}

func TestGlobalSearch(t *testing.T) {
	resp, err := client.GlobalSearch(context.Background(), &pb.GlobalSearchRequest{
		SearchTerm: "Маски",
	})
	if err != nil {
		t.Fatalf("GlobalSearch failed: %v", err)
	}

	t.Logf("GlobalSearch response: %+v", resp)
	assert.Greater(t, len(resp.Entry), 0)
}

func TestTrackEntry(t *testing.T) {
	// to ensure state is clear from previous tests
	client.DeleteAllTracked(context.Background(), &pb.DeleteAllTrackedRequest{})

	resp, err := client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
		Key: &pb.TrackedEntryKey{
			EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
			EntityId:   34145,
			UserId:     "testuser",
		},
		ForceUpdate: false,
	})
	if err != nil {
		t.Fatalf("TrackEntry failed: %v", err)
	}

	assert.Equal(t, resp.Key.EntityId, int64(34145))
	assert.Equal(t, resp.Result, pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK)

	// List tracked entries
	resp2, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{
		UserId: "testuser",
	})
	assert.Equal(t, len(resp2.Entry), 1)

	// Untrack entry now
	resp3, err := client.UntrackEntry(context.Background(), &pb.UntrackEntryRequest{
		Key: &pb.TrackedEntryKey{
			EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
			EntityId:   34145,
			UserId:     "testuser",
		},
	})

	assert.Equal(t, resp3.Result, pb.UntrackEntryResult_UNTRACK_ENTRY_RESULT_OK)
	resp4, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{
		UserId: "testuser",
	})
	assert.Equal(t, len(resp4.Entry), 0)
}

func TestArchiveEntry(t *testing.T) {
	// to ensure state is clear from previous tests
	client.DeleteAllTracked(context.Background(), &pb.DeleteAllTrackedRequest{})

	resp, err := client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
		Key: &pb.TrackedEntryKey{
			EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
			EntityId:   34145,
			UserId:     "testuser",
		},
		ForceUpdate: false,
	})
	if err != nil {
		t.Fatalf("TrackEntry failed: %v", err)
	}

	assert.Equal(t, resp.Key.EntityId, int64(34145))
	assert.Equal(t, resp.Result, pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK)

	// List tracked entries
	resp2, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{
		UserId: "testuser",
	})
	assert.Equal(t, len(resp2.Entry), 1)

	entry := resp2.Entry[0]
	assert.Equal(t, entry.Key.EntityId, int64(34145))

	// Archive entry now
	_, err = client.UpdateEntry(context.Background(), &pb.UpdateTrackedEntryRequest{TrackedEntry: entry})
	assert.Nil(t, err)

	resp3, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{
		UserId: "testuser",
	})
	assert.Equal(t, len(resp2.Entry), 1)

	entry2 := resp2.Entry[0]
	assert.Equal(t, entry, entry2)

}

func TestMain(m *testing.M) {
	if os.Getenv("FLIBUSTIER_INTEGRATION") != "1" {
		// Not running integration tests unless explicitly requested
		os.Exit(0)
	}

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client = pb.NewFlibustierServiceClient(conn)

	ret := m.Run()

	conn.Close()
	os.Exit(ret)
}
