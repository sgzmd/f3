package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/f3/web/rpc"
)

var decoder = schema.NewDecoder()

type TrackPageHandler struct {
	http.Handler

	client rpc.ClientInterface
}

type TrackRequestSchema struct {
	EntryId   int64
	EntryType string
}

func NewTrackPageHandler(client rpc.ClientInterface) *TrackPageHandler {
	return &TrackPageHandler{client: client}
}

func (page *TrackPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Problem parsing form: %+v", err)
		ErrorToBrowser(w, r, err)
	}
	var trackReq TrackRequestSchema
	err = decoder.Decode(&trackReq, r.Form)
	if err != nil {
		log.Printf("Problem decoding form: %+v", err)
		ErrorToBrowser(w, r, err)
	}

	resp, err := page.client.TrackEntry(&proto.TrackEntryRequest{Key: &proto.TrackedEntryKey{
		EntityId:   trackReq.EntryId,
		EntityType: proto.EntryType(proto.EntryType_value[trackReq.EntryType]),
		UserId:     "default",
	}})

	if err != nil {
		log.Printf("Error tracking entry: %+v", err)
		ErrorToBrowser(w, r, err)
	}

	http.Redirect(w, r,
		fmt.Sprintf("/?track_result=%s", resp.Result),
		http.StatusTemporaryRedirect)
}
