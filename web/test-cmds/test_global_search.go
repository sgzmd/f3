package main

import (
	"fmt"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
)

func TryGlobalSearch() {
	conn, err := grpc.Dial("172.23.22.238:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewFlibustierServiceClient(conn)
	search := rpc.NewGrpcSearch(client)
	resp, err := search.GlobalSearch("Маски")

	if err != nil {
		fmt.Errorf("Failed to query GRPC: %s", err)
	} else {
		fmt.Printf("%s", prototext.Format(resp))
	}
}
