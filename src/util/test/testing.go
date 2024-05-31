package utiltesting

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/moura1001/ramengo-api/src/dto"
	"github.com/moura1001/ramengo-api/src/handlers"
	"github.com/moura1001/ramengo-api/src/model"
)

func NewHttpTestServer() *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /broths", handlers.WithRequiredHeaders(handlers.HandleBrothList))

	return httptest.NewServer(router)
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertErrorResponse(t *testing.T, body io.Reader, want dto.ErrorResponse) {
	t.Helper()
	var got dto.ErrorResponse
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse response from server %q into dto.ErrorResponse: '%v'", body, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t *testing.T, response *http.Response, want string) {
	t.Helper()
	if response.Header.Get(handlers.HeaderContentType) != want {
		t.Errorf("response did not have %s of %s, got %v", handlers.HeaderContentType, want, response.Header)
	}
}

func AssertBrothList(t *testing.T, body io.Reader, want []model.Broth) {
	t.Helper()
	var got []model.Broth
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse response from server %q into broth list: '%v'", body, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
