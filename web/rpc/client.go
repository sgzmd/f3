package rpc

import (
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewFlibustierClient(backend string) (*pb.FlibustierServiceClient, error) {
	conn, err := grpc.Dial(backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewFlibustierServiceClient(conn)
	return &client, nil
}
