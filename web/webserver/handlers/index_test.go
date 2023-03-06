package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ericchiang/css"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
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

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	doc, err := html.Parse(resp.Body)

	assert.Nil(t, err)

	sel, err := css.Parse(`a[data-testid="tracked-entry"]`)
	assert.Nil(t, err)

	numFound := 0
	for i, el := range sel.Select(doc) {
		found := false
		for _, attr := range el.Attr {
			if attr.Key == "data-testval" {
				assert.Equal(t, EXPECTED_IDS[i], attr.Val)
				found = true
				numFound++
			}
		}
		assert.True(t, found)
	}

	assert.Equal(t, len(EXPECTED_IDS), numFound)
}
