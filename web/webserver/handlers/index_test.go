package handlers

import (
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Tests IndexHandler
func TestIndexHandler(t *testing.T) {
	app, clientContext := CreateTestingApp()

	EXPECTED_IDS := []string{"79232", "109170"}

	app.Get("/", IndexHandler(clientContext))

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:    tgauth.DefaultCookieName,
		Value:   "something",
		Path:    "/",
		Expires: time.Now().Add(time.Hour + 24),
	})

	resp, err := app.Test(req, -1)
	log.Printf("Response: %+v", resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	divFound := false

	h, _ := doc.Html()
	log.Printf("HTML: %s", h)

	doc.Find(`div[data-testid="tracked-entries-dev"]`).Each(func(i int, s *goquery.Selection) {
		divFound = true
	})

	numProcessedLinks := 0

	doc.Find(`*[data-testid="tracked-entry"`).Each(func(i int, link *goquery.Selection) {
		id, err := link.Attr("data-testval")
		assert.Nil(t, err)
		assert.Equal(t, EXPECTED_IDS[i], id)
		numProcessedLinks++
	})

	assert.Equal(t, len(EXPECTED_IDS), numProcessedLinks)

	assert.True(t, divFound)
}
