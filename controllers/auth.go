package controllers

import (
	"github.com/gorilla/mux"
	"github.com/manhtai/cusbot/config"

	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

// Index is the index page
func Index(w http.ResponseWriter, r *http.Request) {
	config.Templ.ExecuteTemplate(w, "index.html", nil)
}

// Login servers our login page
func Login(w http.ResponseWriter, r *http.Request) {
	config.Templ.ExecuteTemplate(w, "login.html", nil)
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("name")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/auth/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// MustAuth is a decorator for login required url
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// LoginHandle redirect user to providers' login page & receive callback from them
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	action := vars["action"]
	provider := vars["provider"]

	switch action {
	case "login":
		q := r.URL.Query()
		q.Set("provider", provider)

		r.URL.RawQuery = q.Encode()

		// FIXME: Find a way to change callback url in place
		config.CreateProvider("https://" + r.Host + "/auth/callback/" + provider)
		gothic.BeginAuthHandler(w, r)

	case "callback":
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		userName := base64.StdEncoding.EncodeToString([]byte(user.Name))
		http.SetCookie(
			w,
			&http.Cookie{
				Name:  "name",
				Value: userName,
				Path:  "/",
			})
		w.Header().Set("Location", "/channel")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

// Logout uses to log User out
func Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Log user out
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
