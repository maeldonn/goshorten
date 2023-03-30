package api

import (
	"net/http"
	"strings"

	"github.com/maeldonn/url-shortener/database"
	"github.com/maeldonn/url-shortener/types"
)

type URLHandler struct{}

func NewURLHandler() *URLHandler {
	return &URLHandler{}
}

func (*URLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetMethod(w, r)
	case http.MethodPost:
		handlePostMethod(w, r)
	case http.MethodDelete:
		handleDeleteMethod(w, r)
	default:
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}

func handleGetMethod(w http.ResponseWriter, r *http.Request) {
	slug := strings.Split(r.URL.Path, "/")[1]
	url, err := database.GetURL(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func handlePostMethod(w http.ResponseWriter, r *http.Request) {
	shortened := types.ExtractShortenedURL(r.Body)
	if err := database.SaveSlug(shortened); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func handleDeleteMethod(w http.ResponseWriter, r *http.Request) {
	slug := strings.Split(r.URL.Path, "/")[1]
	database.RemoveSlug(slug)
}
