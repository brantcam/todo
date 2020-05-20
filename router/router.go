package router

import (
	"net/http"
	"todo/handlers"
	"todo/todo"

	"github.com/gorilla/mux"
)

// Opts will contain all of the handler functions for each type
type Opts struct {
	Todo todo.Repo
}

// New creates a new instance of a gorilla mux router
func New(opts Opts) *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/").Handler(handlers.GetList(opts.Todo))

	return r
}
