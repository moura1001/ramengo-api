package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/service"
)

func HandleOrderNew(w http.ResponseWriter, r *http.Request) {
	var orderRequest dto.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil || !isValidOrderRequest(orderRequest) {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse("both brothId and proteinId are required"))
		return
	}

	orderResponse, code, err := service.GetOrderProcessor().ProcessOrder(orderRequest)
	if err != nil {
		WriteJSON(w, code, dto.NewErrorResponse(err.Error()))
		return
	}

	WriteJSON(w, code, orderResponse)
}

func isValidOrderRequest(orderRequest dto.OrderRequest) bool {
	isValidBrothId := len(strings.TrimSpace(orderRequest.BrothId)) > 0
	isValidProteinId := len(strings.TrimSpace(orderRequest.ProteinId)) > 0

	return isValidBrothId && isValidProteinId
}
