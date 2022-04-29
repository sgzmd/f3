package main

import (
	"context"
	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func dialer(flibustaDb string) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	srv, err := NewServer(flibustaDb, "" /* in memory datastore */)
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
