package handlers

import (
	"encoding/json"
	"io/ioutil"
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

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, err)
		}

		if err := json.Unmarshal(body, &items); err != nil {
			writeJSONResponse(w, http.StatusInternalServerError, err)
		}

		listItems, err := t.AddItem(items)
		if err != nil {
			writeJSONResponse(w, http.StatusServiceUnavailable, err)
		}

		writeJSONResponse(w, http.StatusOK, listItems)
	}
}
