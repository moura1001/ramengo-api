package handlers

import (
	"net/http"

	"github.com/moura1001/ramengo-api/src/dto"
)

func WithRequiredHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") == "" {
			WriteJSON(w, http.StatusForbidden, dto.NewErrorResponse("x-api-key header missing"))
			return
		}

		handler.ServeHTTP(w, r)
	}
}
