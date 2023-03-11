package integration

import (
	"context"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
	"io"
	"net/http"
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

func TestGlobalSearch_Author(t *testing.T) {
	resp, err := client.GlobalSearch(context.Background(), &pb.GlobalSearchRequest{
		SearchTerm: "Сергей Мусаниф",
	})
	if err != nil {
		t.Fatalf("GlobalSearch failed: %v", err)
	}

	t.Logf("GlobalSearch response: %+v", resp)
	assert.Equal(t, len(resp.Entry), 1)
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

func TestArchiveEntryBackend(t *testing.T) {
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
	assert.Equal(t, len(resp3.Entry), 1)

	entry2 := resp3.Entry[0]
	text1 := prototext.Format(entry)
	text2 := prototext.Format(entry2)
	assert.Equal(t, text1, text2)

}

func TestForceRefresh(t *testing.T) {
	result, err := client.ForceRefresh(context.Background(), &pb.ForceRefreshRequest{})
	assert.Nil(t, err)
	assert.Equal(t, result.Result, proto.ForceRefreshResponse_FORCE_REFRESH_RESULT_OK)

	// Running test again to ensure we can still run it after refresh
	TestGlobalSearch(t)
}

func TestSmokeTestWeb(t *testing.T) {
	s := getHtml(t, "http://localhost:8088")
	assert.Containsf(t, s, "<h4>Отслеживаем</h4>", "Expected to find 'Отслеживаем' in response, got: %s", s)
}

func DisabledTestArchiveEntryWeb(t *testing.T) {
	// to ensure state is clear from previous tests
	client.DeleteAllTracked(context.Background(), &pb.DeleteAllTrackedRequest{})

	// Ignoring errors because this is tested elsewhere
	_, _ = client.TrackEntry(context.Background(), &pb.TrackEntryRequest{
		Key: &pb.TrackedEntryKey{
			EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
			EntityId:   34145,
			UserId:     "testuser",
		},
		ForceUpdate: false,
	})

	index := getHtml(t, "http://localhost:8088")
	assert.Containsf(t, index,
		`<a target="_blank" href="http://flibusta.is/s/34145" data-testid="tracked-entry" data-testval="34145">Маски [= Унесенный ветром]</a>`,
		"Expected to find entry 34145 in response, got: %s", index)

	// Ignoring body because redirect anyway
	_, e := http.Get("http://localhost:8088/archive/1/34145")
	assert.Nil(t, e, "Failed to archive entry: %+v", e)

	archives := getHtml(t, "http://localhost:8088/archives")
	assert.Containsf(t, archives,
		`<a target="_blank" href="http://flibusta.is/s/34145" data-testid="tracked-entry" data-testval="34145">Маски [= Унесенный ветром]</a>`,
		"Expected to find entry 34145 in response, got: %s", index)
}

func getHtml(t *testing.T, url string) string {
	resp, err := http.Get(url)
	assert.Nil(t, err, "Failed to get URL %s: %+v", url, err)

	d, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, "Failed to read response body: %+v", err)

	return string(d)
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
