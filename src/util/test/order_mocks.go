package utiltesting

import (
	"fmt"
	"net/http"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/service"
)

func OrderProcessorInternalServerErrorMock(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	return dto.OrderResponse{}, http.StatusInternalServerError, fmt.Errorf("could not place order")
}

func OrderProcessorCreated(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	return OrderResponseSuccessfully, http.StatusCreated, nil
}

func GetOrderProcessorMock(mockType int) service.OrderProcessor {
	switch mockType {
	case http.StatusInternalServerError:
		return service.OrderProcessorFunc(OrderProcessorInternalServerErrorMock)
	case http.StatusCreated:
		return service.OrderProcessorFunc(OrderProcessorCreated)
	default:
		return service.OrderProcessorFunc(OrderProcessorInternalServerErrorMock)
	}
}
