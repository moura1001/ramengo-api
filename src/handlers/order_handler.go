package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/service"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

func HandleOrderNew(w http.ResponseWriter, r *http.Request) {
	var orderRequest dto.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil || !isValidOrderRequest(orderRequest) {
		WriteJSON(w, http.StatusBadRequest, dto.NewErrorResponse("both brothId and proteinId are required"))
		return
	}

	ctx := context.WithValue(context.Background(), utilapp.HeaderXApiKey, r.Header.Get(utilapp.HeaderXApiKey))

	orderResponse, code, err := service.GetOrderProcessor().ProcessOrder(ctx, orderRequest)
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
