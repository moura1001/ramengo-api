package handlers

import (
	"net/http"

	"github.com/moura1001/ramengo-api/src/model"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

func HandleProteinList(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, model.ListAllProteins())
}

func HandleProteinOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utilapp.HeaderAccessControlAllowOrigin, utilapp.AllowedHttpOrigin)
	w.Header().Set(utilapp.HeaderAccessControlAllowMethods, "GET, OPTIONS")
	w.Header().Set(utilapp.HeaderAccessControlAllowHeaders, utilapp.AllowedHeaders)
	w.Header().Set(utilapp.HeaderAccessControlMaxAge, "86400")
	w.WriteHeader(http.StatusNoContent)
}
