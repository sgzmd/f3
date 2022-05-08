package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
)

type SearchResultType struct {
	Entry []*proto.FoundEntry
}

type IndexPage struct {
	DefaultSearchTerm string
	TrackResult       proto.TrackEntryResult
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
	log.Printf("ServeHTTP: %s; query=%s", r.RequestURI, r.URL.Query())
	searchTerm, ok := r.URL.Query()["searchTerm"]

	data := IndexPage{
		SearchResult: idx.getSearchResults(r),
		TrackResult:  0,
	}

	if ok {
		idx.searchTerm = searchTerm[0]
	} else {
		idx.searchTerm = ""
	}

	data.DefaultSearchTerm = idx.searchTerm
	tracked, ok := r.URL.Query()["track_result"]

	if ok {
		r, e := strconv.Atoi(tracked[0])
		if e != nil {
			log.Printf("Failed to parse tracked status: %+v", e)
		} else {
			data.TrackResult = proto.TrackEntryResult(r)
		}
	}

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
