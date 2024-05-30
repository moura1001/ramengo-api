package handlers

import (
	"encoding/json"
	"net/http"
)

func WithRequiredHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "x-api-key header missing",
			})
			return
		}

		handler.ServeHTTP(w, r)
	}
}
