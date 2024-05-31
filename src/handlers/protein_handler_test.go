package handlers_test

import (
	"net/http"
	"testing"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/handlers"
	"github.com/moura1001/ramengo-api/src/model"
	utiltesting "github.com/moura1001/ramengo-api/src/util/test"
)

func TestProteinEndpoint(t *testing.T) {
	server := utiltesting.NewHttpTestServer()
	defer server.Close()
	client := server.Client()

	t.Run("should return 403 forbidden when the request does not have the x-api-key header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL+"/proteins", nil)
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusForbidden)
		utiltesting.AssertContentType(t, response, handlers.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("x-api-key header missing"))
	})

	t.Run("should return 200 with the list of available proteins", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL+"/proteins", nil)
		request.Header.Set("x-api-key", "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		utiltesting.AssertContentType(t, response, handlers.JsonContentType)
		utiltesting.AssertProteinList(t, response.Body, model.ListAllProteins())
	})
}
