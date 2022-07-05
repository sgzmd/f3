package handlers

import (
	"fmt"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

type SearchResultType struct {
	Entry []*proto.FoundEntry
}

type Entry struct {
	Entry *proto.TrackedEntry
}

type TrackedEntriesType struct {
	Entry []Entry
}

type IndexPage struct {
	DefaultSearchTerm string
	TrackResult       proto.TrackEntryResult
	SearchResult      *SearchResultType
	TrackedEntries    *TrackedEntriesType
}

type IndexPageHandler struct {
	http.Handler

	client     rpc.ClientInterface
	searchTerm string
	auth       tgauth.TelegramAuth
}

func NewIndexPageHandler(client rpc.ClientInterface, auth tgauth.TelegramAuth) *IndexPageHandler {
	return &IndexPageHandler{
		client: client,
		auth:   auth,
	}
}

func (e *Entry) EntryUrl() string {
	return TrackedEntryUrl(e.Entry)
}

func (idx *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServeHTTP: %s; query=%s", r.RequestURI, r.URL.Query())

	params, err := idx.auth.GetParamsFromCookie(r)
	if err != nil {
		log.Printf("Unable to get params from cookie: %+v", err)
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	ok, err := idx.auth.CheckAuth(params)
	if err != nil {
		log.Printf("Unable to check auth: %+v", err)
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	} else if !ok {
		log.Printf("Auth is not ok")
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	userInfo, err := idx.auth.GetUserInfo(params)
	if err != nil {
		log.Printf("Auth is not ok")
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	searchTerm, ok := r.URL.Query()["searchTerm"]

	if ok {
		idx.searchTerm = searchTerm[0]
	} else {
		idx.searchTerm = ""
	}

	data := IndexPage{
		SearchResult:   idx.getSearchResults(r),
		TrackedEntries: idx.getTrackedEntries(w, r, userInfo),
		TrackResult:    0,
	}

	data.DefaultSearchTerm = idx.searchTerm
	tracked, ok := r.URL.Query()["track_result"]

	if ok {
		data.TrackResult = proto.TrackEntryResult(proto.TrackEntryResult_value[tracked[0]])
	}

	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(filename)

	t := template.Must(template.ParseFiles(dir + "/../templates/index.html"))
	err = t.Execute(w, data)
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

func (idx *IndexPageHandler) getTrackedEntries(w http.ResponseWriter, r *http.Request, info *tgauth.UserInfo) *TrackedEntriesType {
	resp, err := idx.client.ListTrackedEntries(&proto.ListTrackedEntriesRequest{UserId: MakeUserKey(info)})
	if err != nil {
		log.Printf("Error fetching tracked entries: %+v", err)
		ErrorToBrowser(w, r, err)
	}

	entries := make([]Entry, len(resp.Entry))
	for i := 0; i < len(resp.Entry); i++ {
		entries[i] = Entry{
			Entry: resp.Entry[i],
		}
	}

	return &TrackedEntriesType{Entry: entries}
}
