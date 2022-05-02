// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	"flag"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	handlers "github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/f3/web/rpc"
	"log"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
)

const (
	StaticPrefix = "/static/"
)

var (
	useFakes    *bool
	grpcBackend *string
)

func main() {
	useFakes = flag.Bool("use_fake_search", false, "Whether use fake search or real GRPC")
	grpcBackend = flag.String("grpc_backend", "", "GRPC backend to use if any")
	flag.Parse()

	var search rpc.Search
	var grpcClient *pb.FlibustierServiceClient

	if *useFakes {
		s := rpc.FakeSearch{}
		search = s.Search
	} else {
		var err error
		grpcClient, err = rpc.NewFlibustierClient(*grpcBackend)
		if err != nil {
			panic(err)
		}

		search = rpc.NewGrpcSearch(*grpcClient)
		log.Printf("Using GRPC backend %s", *grpcBackend)
	}

	r := mux.NewRouter()

	r.Handle("/", handlers.NewIndexPageHandler(search)).Methods("GET")

	r.PathPrefix(StaticPrefix).Handler(http.StripPrefix(StaticPrefix, http.FileServer(http.Dir("./templates/"+StaticPrefix))))

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", r)
}
