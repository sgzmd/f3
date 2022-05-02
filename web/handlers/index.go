package handlers

import (
	"fmt"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
	"html/template"
	"net/http"
)

type SearchResultType struct {
	Entry []*proto.FoundEntry
}

type IndexPage struct {
	DefaultSearchTerm string
	SearchResult      *SearchResultType
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := IndexPage{DefaultSearchTerm: "Дозор", SearchResult: getSearchResults(r)}
	t := template.Must(template.ParseFiles("./templates/index.html"))
	err := t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
}

func getSearchResults(r *http.Request) *SearchResultType {
	searchTerm, ok := r.URL.Query()["searchTerm"]
	if ok {
		client := rpc.FakeSearch{}
		searchResult, err := client.GlobalSearch(searchTerm[0])
		if err != nil {
			fmt.Errorf("Error querying GRPC: %s", err)
			return nil
		} else {
			return &SearchResultType{Entry: searchResult.Entry}
		}
	} else {
		return nil
	}
}
