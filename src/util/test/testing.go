package utiltesting

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/moura1001/ramengo-api/src/handlers"
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

func AssertErrorResponse(t *testing.T, body io.Reader, want any) {
	t.Helper()
	got := make(map[string]string, 0)
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse response from server %q into map: '%v'", body, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
