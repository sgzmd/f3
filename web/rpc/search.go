package rpc

import (
	"context"
	"github.com/golang/protobuf/proto"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
)

type Search interface {
	GlobalSearch(searchTerm string) (*pb.GlobalSearchResponse, error)
}
type FakeSearch struct {
	Search
}

type GrpcSearch struct {
	Search
	client pb.FlibustierServiceClient
}

func NewGrpcSearch(client pb.FlibustierServiceClient) *GrpcSearch {
	return &GrpcSearch{client: client}
}

func (fs *FakeSearch) GlobalSearch(searchTerm string) (*pb.GlobalSearchResponse, error) {

	resp := pb.GlobalSearchResponse{}
	proto.UnmarshalText(FakeResponse, &resp)
	return &resp, nil
	//
	//entries := make([]*pb.FoundEntry, 1)
	//entries[0] = &pb.FoundEntry{
	//	EntryId:     123,
	//	EntryName:   searchTerm,
	//	EntryType:   pb.EntryType_ENTRY_TYPE_AUTHOR,
	//	Author:      searchTerm + " author",
	//	NumEntities: 3}
	//resp := pb.GlobalSearchResponse{
	//	OriginalRequest: &pb.GlobalSearchRequest{
	//		SearchTerm:      searchTerm,
	//		EntryTypeFilter: pb.EntryType_ENTRY_TYPE_UNSPECIFIED,
	//	},
	//	Entry: entries,
	//}
	//
	//return &resp, nil
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
