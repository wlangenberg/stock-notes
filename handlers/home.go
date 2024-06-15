package handlers

import (
	"net/http"
	"stocknotes/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
