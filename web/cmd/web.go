// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"

	handlers "github.com/sgzmd/f3/web/handlers"
	"github.com/sgzmd/f3/web/rpc"
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

	auth tgauth.TelegramAuth
)

type Options struct {
	GrpcBackend string `short:"g" long:"grpc_backend" description:"GRPC Backend to use"`
	//TelegramToken string `short:"t" long:"telegram_token" description:"Telegram token to use" required:"true"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)

	client, err := rpc.NewClient(&opts.GrpcBackend)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.Handle("/", handlers.NewIndexPageHandler(*client)).Methods("GET")
	r.Handle("/track", handlers.NewTrackPageHandler(*client)).Methods("GET")

	r.PathPrefix(StaticPrefix).Handler(http.StripPrefix(StaticPrefix, http.FileServer(http.Dir("./templates/"+StaticPrefix))))

	e := http.ListenAndServe(":8080", r)
	log.Fatal(e)
}
