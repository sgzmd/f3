package rpc

import (
	"log"

	"google.golang.org/protobuf/encoding/prototext"

	"github.com/golang/protobuf/proto"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientInterface interface {
	GlobalSearch(in *pb.GlobalSearchRequest) (*pb.GlobalSearchResponse, error)
	CheckUpdates(in *pb.CheckUpdatesRequest) (*pb.CheckUpdatesResponse, error)
	GetSeriesBooks(in *pb.GetSeriesBooksRequest) (*pb.GetSeriesBooksResponse, error)
	GetAuthorBooks(in *pb.GetAuthorBooksRequest) (*pb.GetAuthorBooksResponse, error)
	TrackEntry(in *pb.TrackEntryRequest) (*pb.TrackEntryResponse, error)
	ListTrackedEntries(in *pb.ListTrackedEntriesRequest) (*pb.ListTrackedEntriesResponse, error)
	UntrackEntry(in *pb.UntrackEntryRequest) (*pb.UntrackEntryResponse, error)
	UpdateEntry(in *pb.UpdateTrackedEntryRequest) (*pb.UpdateTrackedEntryResponse, error)
	GetUserInfo(in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error)
	ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
}

type FakeClientImplementation struct {
}

func (f FakeClientImplementation) ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	resp := pb.ListUsersResponse{}
	return &resp, nil
}

func (f FakeClientImplementation) GetUserInfo(in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return &pb.GetUserInfoResponse{}, nil
}

type GrpcClientImplementation struct {
	client pb.FlibustierServiceClient
}

func (g GrpcClientImplementation) ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return g.client.ListUsers(context.Background(), in)
}

func (g GrpcClientImplementation) GetUserInfo(in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	log.Printf("GetUserInfo: %+v", in)
	return g.client.GetUserInfo(context.Background(), in)
}

func (g GrpcClientImplementation) GlobalSearch(in *pb.GlobalSearchRequest) (*pb.GlobalSearchResponse, error) {
	log.Printf("GlobalSearch: %+v", in)
	return g.client.GlobalSearch(context.Background(), in)
}

func (g GrpcClientImplementation) CheckUpdates(in *pb.CheckUpdatesRequest) (*pb.CheckUpdatesResponse, error) {
	return g.client.CheckUpdates(context.Background(), in)
}

func (g GrpcClientImplementation) GetSeriesBooks(in *pb.GetSeriesBooksRequest) (*pb.GetSeriesBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcClientImplementation) GetAuthorBooks(in *pb.GetAuthorBooksRequest) (*pb.GetAuthorBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcClientImplementation) TrackEntry(in *pb.TrackEntryRequest) (*pb.TrackEntryResponse, error) {
	return g.client.TrackEntry(context.Background(), in)
}

func (g GrpcClientImplementation) ListTrackedEntries(in *pb.ListTrackedEntriesRequest) (*pb.ListTrackedEntriesResponse, error) {
	log.Printf("ListTrackedEntries: %+v", in)
	return g.client.ListTrackedEntries(context.Background(), in)
}

func (g GrpcClientImplementation) UntrackEntry(in *pb.UntrackEntryRequest) (*pb.UntrackEntryResponse, error) {
	return g.client.UntrackEntry(context.Background(), in)
}

func (g GrpcClientImplementation) UpdateEntry(in *pb.UpdateTrackedEntryRequest) (*pb.UpdateTrackedEntryResponse, error) {
	return g.client.UpdateEntry(context.Background(), in)
}

func (f FakeClientImplementation) UpdateEntry(in *pb.UpdateTrackedEntryRequest) (*pb.UpdateTrackedEntryResponse, error) {
	return &pb.UpdateTrackedEntryResponse{TrackedEntry: in.TrackedEntry}, nil
}

func (f FakeClientImplementation) GlobalSearch(in *pb.GlobalSearchRequest) (*pb.GlobalSearchResponse, error) {
	resp := pb.GlobalSearchResponse{}
	proto.UnmarshalText(GlobalSearchFakeResponse, &resp)
	return &resp, nil
}

func (f FakeClientImplementation) CheckUpdates(in *pb.CheckUpdatesRequest) (*pb.CheckUpdatesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FakeClientImplementation) GetSeriesBooks(in *pb.GetSeriesBooksRequest) (*pb.GetSeriesBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FakeClientImplementation) GetAuthorBooks(in *pb.GetAuthorBooksRequest) (*pb.GetAuthorBooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FakeClientImplementation) TrackEntry(in *pb.TrackEntryRequest) (*pb.TrackEntryResponse, error) {
	return &pb.TrackEntryResponse{Key: in.Key, Result: pb.TrackEntryResult_TRACK_ENTRY_RESULT_OK}, nil
}

func (f FakeClientImplementation) ListTrackedEntries(in *pb.ListTrackedEntriesRequest) (*pb.ListTrackedEntriesResponse, error) {
	resp := pb.ListTrackedEntriesResponse{}
	err := prototext.Unmarshal([]byte(ListEntriesFakeResponse), &resp)

	if err != nil {
		log.Fatalf("Failed to unmarshal ListTrackedEntriesResponse: %v", err)
	}
	return &resp, nil
}

func (f FakeClientImplementation) UntrackEntry(in *pb.UntrackEntryRequest) (*pb.UntrackEntryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewClient(backend *string) (ClientInterface, error) {
	var client ClientInterface
	if backend == nil || *backend == "" {
		log.Print("Using FakeClientImplementation")
		client = FakeClientImplementation{}
	} else {
		log.Printf("Attempting to use GrpcClientImplementation with backend=%s", *backend)
		conn, err := grpc.Dial(*backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		fclient := pb.NewFlibustierServiceClient(conn)
		client = GrpcClientImplementation{client: fclient}
	}
	return client, nil
}

func NewFlibustierClient(backend string) (*pb.FlibustierServiceClient, error) {
	conn, err := grpc.Dial(backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewFlibustierServiceClient(conn)
	return &client, nil
}
