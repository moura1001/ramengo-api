package handlers

import (
	"net/http"

	"github.com/moura1001/ramengo-api/src/dto"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

func WithRequiredHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(utilapp.HeaderXApiKey) == "" {
			WriteJSON(w, http.StatusForbidden, dto.NewErrorResponse("x-api-key header missing"))
			return
		}

		w.Header().Set(utilapp.HeaderAccessControlAllowOrigin, utilapp.AllowedHttpOrigin)

		handler.ServeHTTP(w, r)
	}
}
