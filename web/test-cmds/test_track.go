package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
)

func TryTrack() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewFlibustierServiceClient(conn)

	resp, err := client.TrackEntry(context.Background(), &pb.TrackEntryRequest{Key: &pb.TrackedEntryKey{
		EntityId:   34145,
		EntityType: pb.EntryType_ENTRY_TYPE_SERIES,
		UserId:     "user",
	}})

	if err != nil {
		log.Fatalf("Failed to query GRPC: %s", err)
	} else {
		log.Printf("%s", prototext.Format(resp))
	}

	resp2, err := client.ListTrackedEntries(context.Background(), &pb.ListTrackedEntriesRequest{UserId: "user"})

	if err != nil {
		log.Fatalf("Failed to query GRPC: %s", err)
	} else {
		fmt.Printf("%s", prototext.Format(resp2))
	}
}
