package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	HeaderContentType = "Content-Type"
	JsonContentType   = "application/json"
)

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set(HeaderContentType, JsonContentType)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
