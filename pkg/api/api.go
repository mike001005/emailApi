package api

import (
	"github.com/gorilla/mux"
)

// Handlers registers a new route with a matcher for the URL path
func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/send", Send)

	return r
}
