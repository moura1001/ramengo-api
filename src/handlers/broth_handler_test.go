package handlers_test

import (
	"net/http"
	"testing"

	utiltesting "github.com/moura1001/ramengo-api/src/util/test"
)

func TestBrothEndpoint(t *testing.T) {
	server := utiltesting.NewHttpTestServer()
	defer server.Close()
	client := server.Client()

	t.Run("should return 403 forbidden when the request does not have the x-api-key header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, server.URL+"/broths", nil)
		response, _ := client.Do(request)

		utiltesting.AssertStatus(t, response.StatusCode, http.StatusForbidden)
		utiltesting.AssertErrorResponse(t, response.Body, map[string]string{
			"error": "x-api-key header missing",
		})
	})
}
