package controllers

import (
	"net/http"

	"github.com/manhtai/cusbot/config"
)

// IndexView is the index page
func Index(w http.ResponseWriter, r *http.Request) {
	config.Templ.ExecuteTemplate(w, "index.html", nil)
}
