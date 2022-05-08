package handlers

import (
	"net/http"

	"github.com/sgzmd/f3/web/rpc"
)

type TrackPageHandler struct {
	http.Handler

	client rpc.ClientInterface
}

func NewTrackPageHandler(client rpc.ClientInterface) *TrackPageHandler {
	return &TrackPageHandler{client: client}
}

func (idx *TrackPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/?track_result=true", http.StatusTemporaryRedirect)
}
