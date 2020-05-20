package handlers

import (
	"encoding/json"
	"net/http"
	"todo/list"
)

// GetList will return all items on the todo list
func GetList(t list.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listItems, err := t.GetList()
		if err != nil {
			writeJSONResponse(w, http.StatusServiceUnavailable, err)
		}
		writeJSONResponse(w, http.StatusOK, listItems)
	}
}

// AddItem ...
func AddItem(t list.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var items []list.Item
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&items); err != nil {
			writeJSONResponse(w, http.StatusBadRequest, err)
		}
		listItems, err := t.AddItem(items)
		if err != nil {
			writeJSONResponse(w, http.StatusServiceUnavailable, err)
		}
		writeJSONResponse(w, http.StatusOK, listItems)
	}
}
