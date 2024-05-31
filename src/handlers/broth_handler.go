package handlers

import (
	"net/http"

	"github.com/moura1001/ramengo-api/src/model"
)

func HandleBrothList(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, model.ListAllBroths())
}
