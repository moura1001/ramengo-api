package handlers

import (
	"encoding/json"
	"net/http"

	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set(utilapp.HeaderContentType, utilapp.JsonContentType)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
