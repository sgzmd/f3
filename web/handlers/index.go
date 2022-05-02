package handlers

import (
	"fmt"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
	"html/template"
	"log"
	"net/http"
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

	search     rpc.Backend
	searchTerm string
}

func NewIndexPageHandler(search rpc.Backend) *IndexPageHandler {
	return &IndexPageHandler{
		search: search,
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

	searchResult, err := idx.search.GlobalSearch(idx.searchTerm)
	if err != nil {
		fmt.Errorf("Error querying GRPC: %s", err)
		return nil
	} else {
		return &SearchResultType{Entry: searchResult.Entry}
	}
}
