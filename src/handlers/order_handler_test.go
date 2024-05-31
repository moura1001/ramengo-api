package handlers_test

import (
	"net/http"
	"testing"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/handlers"
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
		utiltesting.AssertContentType(t, response, handlers.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("x-api-key header missing"))
	})

	t.Run("should return 400 invalid request when the request does not have the requiered body", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, server.URL+"/orders", nil)
		request.Header.Set("x-api-key", "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusBadRequest)
		utiltesting.AssertContentType(t, response, handlers.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("both brothId and proteinId are required"))
	})
}
