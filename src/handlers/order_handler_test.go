package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/service"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
	utiltesting "github.com/moura1001/ramengo-api/src/util/test"
)

func TestOrderEndpoint(t *testing.T) {
	server := utiltesting.NewHttpTestServer()
	defer server.Close()
	client := server.Client()

	t.Run("should return 403 forbidden when the request does not have the x-api-key header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, server.URL+"/orders", nil)
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusForbidden)
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("x-api-key header missing"))
	})

	t.Run("should return 400 invalid request when the request does not have the requiered body", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, server.URL+"/orders", nil)
		request.Header.Set(utilapp.HeaderXApiKey, "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusBadRequest)
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("both brothId and proteinId are required"))
	})

	t.Run("should return 500 internal server error when the external order microservice is down", func(t *testing.T) {
		service.SetOrderProcessor(utiltesting.GetOrderProcessorMock(http.StatusInternalServerError))

		body := dto.OrderRequest{
			BrothId:   "1",
			ProteinId: "1",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(body)
		request, _ := http.NewRequest(http.MethodPost, server.URL+"/orders", payloadBuf)
		request.Header.Set(utilapp.HeaderXApiKey, "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusInternalServerError)
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("could not place order"))
	})

	t.Run("should return 201 when order placed successfully", func(t *testing.T) {
		service.SetOrderProcessor(utiltesting.GetOrderProcessorMock(http.StatusCreated))

		body := dto.OrderRequest{
			BrothId:   "1",
			ProteinId: "1",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(body)
		request, _ := http.NewRequest(http.MethodPost, server.URL+"/orders", payloadBuf)
		request.Header.Set(utilapp.HeaderXApiKey, "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusCreated)
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertOrderResponse(t, response.Body, utiltesting.OrderResponseSuccessfully)
	})

	t.Run("should return 204 determining that the request is safe to send", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodOptions, server.URL+"/orders", nil)
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusNoContent)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowOrigin, utilapp.AllowedHttpOrigin)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowMethods, "POST, OPTIONS")
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowHeaders, utilapp.AllowedHeaders)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlMaxAge, "86400")
	})
}
