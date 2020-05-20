package router

import (
	"net/http"
	"todo/handlers"
	"todo/list"

	"github.com/gorilla/mux"
)

// Opts will contain all of the handler functions for each type
type Opts struct {
	List list.Repo
}

// New creates a new instance of a gorilla mux router
func New(opts Opts) *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/").Handler(handlers.GetList(opts.List))

	return r
}
