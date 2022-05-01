// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	handlers "github.com/sgzmd/f3/web/handlers"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"
)

const (
	StaticPrefix = "/static/"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.PathPrefix(StaticPrefix).Handler(http.StripPrefix(StaticPrefix, http.FileServer(http.Dir("./templates/"+StaticPrefix))))

	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", r)
}
