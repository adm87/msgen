{{- template "disclaimer.go.noedit" }}
package web

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err *ServerError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	jsonData, e := err.Json()

	if e != nil {
		http.Error(w, `{"message":"Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, `{"message":"Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
