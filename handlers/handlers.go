package handlers

import (
	"encoding/json"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, code int, body interface{}) {
	buf, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(buf)
}
