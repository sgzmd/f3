package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
)

type SearchResultType struct {
	Entry []*proto.FoundEntry
}

type IndexPage struct {
	DefaultSearchTerm string
	SearchResult      *SearchResultType
}

type IndexPageHandler struct {
	http.Handler

	client     rpc.ClientInterface
	searchTerm string
}

func NewIndexPageHandler(client rpc.ClientInterface) *IndexPageHandler {
	return &IndexPageHandler{
		client: client,
	}
}

func (idx *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServeHTTP: %s", r.RequestURI)
	searchTerm, ok := r.URL.Query()["searchTerm"]
	data := IndexPage{SearchResult: idx.getSearchResults(r)}

	if ok {
		idx.searchTerm = searchTerm[0]
	} else {
		idx.searchTerm = ""
	}

	data.DefaultSearchTerm = idx.searchTerm

	t := template.Must(template.ParseFiles("./templates/index.html"))
	err := t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
}

func (idx *IndexPageHandler) getSearchResults(r *http.Request) *SearchResultType {
	if idx.searchTerm == "" {
		return nil
	}

	searchResult, err := idx.client.GlobalSearch(&proto.GlobalSearchRequest{
		SearchTerm: idx.searchTerm,
	})

	if err != nil {
		log.Printf("Error querying GRPC: %s", err)
		return nil
	} else {
		return &SearchResultType{Entry: searchResult.Entry}
	}
}
