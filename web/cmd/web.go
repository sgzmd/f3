// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	"flag"
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
	grpcBackend = flag.String("grpc_backend", "", "GRPC backend to use if any")
	flag.Parse()

	client, err := rpc.NewClient(grpcBackend)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.Handle("/", handlers.NewIndexPageHandler(*client)).Methods("GET")

	r.PathPrefix(StaticPrefix).Handler(http.StripPrefix(StaticPrefix, http.FileServer(http.Dir("./templates/"+StaticPrefix))))

	e := http.ListenAndServe(":8080", r)
	log.Fatal(e)
}
