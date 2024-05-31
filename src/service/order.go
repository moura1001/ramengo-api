package service

import (
	"fmt"
	"net/http"

	"github.com/moura1001/ramengo-api/src/dto"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

type OrderProcessor interface {
	ProcessOrder(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error)
}

type OrderProcessorFunc func(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error)

func (p OrderProcessorFunc) ProcessOrder(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	return p(orderRequest)
}

var p OrderProcessor

func GetOrderProcessor() OrderProcessor {
	if utilapp.GetEnv("APP_MODE", "dev") == "prod" {
		return OrderProcessorFunc(orderProcessorDefault)
	}

	if p == nil {
		p = OrderProcessorFunc(orderProcessorDefault)
	}

	return p
}

func SetOrderProcessor(processor OrderProcessor) {
	if processor != nil {
		p = processor
	} else {
		p = OrderProcessorFunc(orderProcessorDefault)
	}
}

func orderProcessorDefault(orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	if orderRequest.BrothId != "1" {
		return dto.OrderResponse{}, http.StatusBadRequest, fmt.Errorf("invalid brothId")
	}

	if orderRequest.ProteinId != "1" {
		return dto.OrderResponse{}, http.StatusBadRequest, fmt.Errorf("invalid proteinId")
	}

	return dto.OrderResponse{
		Id:          "",
		Description: "",
		Image:       "",
	}, http.StatusCreated, nil
}
