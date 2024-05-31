package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/moura1001/ramengo-api/src/dto"
)

func HandleOrderNew(w http.ResponseWriter, r *http.Request) {
	var orderRequest dto.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil || !isValidOrderRequest(orderRequest) {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse("both brothId and proteinId are required"))
		return
	}
}

func isValidOrderRequest(orderRequest dto.OrderRequest) bool {
	isValidBrothId := len(strings.TrimSpace(orderRequest.BrothId)) > 0
	isValidProteinId := len(strings.TrimSpace(orderRequest.ProteinId)) > 0

	return isValidBrothId && isValidProteinId
}
