package handlers

import (
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
)

// LoginHandler HTTP handler for login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(filename)

	t := template.Must(template.ParseFiles(dir + "/../templates/login.html"))
	t.Execute(w, nil)
}

func NewCheckAuthHandler(auth tgauth.TelegramAuth) CheckAuthHandler {
	return CheckAuthHandler{
		auth: auth,
	}
}

type CheckAuthHandler struct {
	http.Handler

	auth tgauth.TelegramAuth
}

func (c CheckAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := make(map[string][]string)
	for k, v := range r.Form {
		params[k] = v
	}

	ok, err := c.auth.CheckAuth(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Invalid auth", http.StatusUnauthorized)
		return
	}

	c.auth.SetCookie(w, params)

	// redirect back to the main page
	http.Redirect(w, r, "/", http.StatusFound)

}
