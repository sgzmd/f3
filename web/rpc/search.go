package rpc

import (
	"context"
	"github.com/golang/protobuf/proto"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
)

type Backend interface {
	GlobalSearch(searchTerm string) (*pb.GlobalSearchResponse, error)
	ListTrackedEntries(userId string) (*pb.ListTrackedEntriesResponse, error)
}
type FakeSearch struct {
	Backend
}

type GrpcSearch struct {
	Backend
	client pb.FlibustierServiceClient
}

func NewGrpcSearch(client pb.FlibustierServiceClient) *GrpcSearch {
	return &GrpcSearch{client: client}
}

func (fs *FakeSearch) GlobalSearch(_ string) (*pb.GlobalSearchResponse, error) {
	resp := pb.GlobalSearchResponse{}
	proto.UnmarshalText(GlobalSearchFakeResponse, &resp)
	return &resp, nil
}

func (gs *GrpcSearch) GlobalSearch(searchTerm string) (*pb.GlobalSearchResponse, error) {
	result, err := gs.client.GlobalSearch(context.Background(), &pb.GlobalSearchRequest{
		SearchTerm:      searchTerm,
		EntryTypeFilter: 0,
	})
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
