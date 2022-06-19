package handlers

import (
	"github.com/sgzmd/f3/web/rpc"
	testing2 "github.com/sgzmd/go-telegram-auth/testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	backend = ""
)

func prepare() (*IndexPageHandler, *httptest.ResponseRecorder) {
	client, _ := rpc.NewClient(&backend)
	idx := NewIndexPageHandler(*client, testing2.NewFakeTelegramAuth(true, "username"))
	rr := httptest.NewRecorder()

	return idx, rr
}

func TestIndexBasic(t *testing.T) {
	idx, rr := prepare()

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	idx.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	str := rr.Body.String()
	assert.True(t, strings.Contains(str, "<h1>Флибустьер!</h1>"))
}

func TestTrackSucceed(t *testing.T) {
	idx, rr := prepare()

	req, err := http.NewRequest("GET", "/?track_result=TRACK_ENTRY_RESULT_OK", nil)
	assert.Nil(t, err)

	idx.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	str := rr.Body.String()
	assert.True(t, strings.Contains(str, `<div id="tracked_status">Добавлено!</div>`))
}
