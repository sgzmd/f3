package main

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"runtime"

	pb "github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

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
	srv, err := NewServerWithDump(flibustaDb, "" /* in memory datastore */, sqlDump)
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
