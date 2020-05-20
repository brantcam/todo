package handlers

import (
	"net/http"
	"todo/todo"
)

// GetList will return all items on the todo list
func GetList(t todo.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listItems, err := t.GetList()
		if err != nil {
			writeJSONResponse(w, http.StatusServiceUnavailable, err)
		}
		writeJSONResponse(w, http.StatusOK, listItems)
	}
}