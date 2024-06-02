package handlers_test

import (
	"net/http"
	"testing"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/model"
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
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
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertErrorResponse(t, response.Body, dto.NewErrorResponse("x-api-key header missing"))
	})

	t.Run("should return 200 with the list of available proteins", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL+"/proteins", nil)
		request.Header.Set(utilapp.HeaderXApiKey, "abc")
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusOK)
		utiltesting.AssertContentType(t, response, utilapp.JsonContentType)
		utiltesting.AssertProteinList(t, response.Body, model.ListAllProteins())
	})

	t.Run("should return 204 determining that the request is safe to send", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodOptions, server.URL+"/proteins", nil)
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusNoContent)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowOrigin, utilapp.AllowedHttpOrigin)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowMethods, "GET, OPTIONS")
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlAllowHeaders, utilapp.AllowedHeaders)
		utiltesting.AssertHeader(t, response, utilapp.HeaderAccessControlMaxAge, "86400")
	})
}
