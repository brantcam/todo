package router

import (
	"net/http"
	"todo/handlers"
	"todo/todo"
)

// Opts will contain all of the handler functions for each type
type Opts struct {
	Todo todo.Repo
}

// New creates a new instance of a router
func New(opts Opts) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", handlers.GetList(opts.Todo))

	return r

}
