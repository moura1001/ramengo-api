package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moura1001/ramengo-api/src/dto"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

type OrderProcessor interface {
	ProcessOrder(ctx context.Context, orderRequest dto.OrderRequest) (dto.OrderResponse, int, error)
}

type OrderProcessorFunc func(ctx context.Context, orderRequest dto.OrderRequest) (dto.OrderResponse, int, error)

func (p OrderProcessorFunc) ProcessOrder(ctx context.Context, orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	return p(ctx, orderRequest)
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

type response struct {
	OrderId string `json:"orderId"`
	err     error
}

func orderProcessorDefault(ctx context.Context, orderRequest dto.OrderRequest) (dto.OrderResponse, int, error) {
	if orderRequest.BrothId != "1" {
		return dto.OrderResponse{}, http.StatusBadRequest, fmt.Errorf("invalid brothId")
	}

	if orderRequest.ProteinId != "1" {
		return dto.OrderResponse{}, http.StatusBadRequest, fmt.Errorf("invalid proteinId")
	}

	orderEndpoint := utilapp.GetEnv("ORDER_ENDPOINT", "")

	xApiKey, _ := ctx.Value(utilapp.HeaderXApiKey).(string)
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	respCh := make(chan response)
	defer cancel()

	go func() {
		req, err := http.NewRequest(http.MethodPost, orderEndpoint, nil)
		if err != nil {
			respCh <- response{err: fmt.Errorf("client: could not create request: %s", err)}
			return
		}
		req.Header.Set(utilapp.HeaderXApiKey, xApiKey)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			respCh <- response{err: fmt.Errorf("client: error making http request: %s", err)}
			return
		}
		if res.StatusCode >= http.StatusBadRequest {
			respCh <- response{err: fmt.Errorf("server: bad response: %s", res.Status)}
			return
		}

		var result response
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			result.err = fmt.Errorf("client: could not read response body: %s", err)
			respCh <- result
			return
		}

		respCh <- result
	}()

	select {
	case <-ctxTimeout.Done():
		return dto.OrderResponse{}, http.StatusInternalServerError, fmt.Errorf("could not place order")
	case resp := <-respCh:
		if resp.err == nil {
			return dto.OrderResponse{
				Id:          resp.OrderId,
				Description: "Salt and Chasu Ramen",
				Image:       "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
			}, http.StatusCreated, nil
		} else {
			return dto.OrderResponse{}, http.StatusInternalServerError, fmt.Errorf("could not place order")
		}
	}
}
