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
	utilapp "github.com/moura1001/ramengo-api/src/util/app"
)

var OrderResponseSuccessfully = dto.OrderResponse{
	Id:          "12345",
	Description: "Salt and Chasu Ramen",
	Image:       "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
}

func NewHttpTestServer() *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /broths", handlers.WithRequiredHeaders(handlers.HandleBrothList))
	router.HandleFunc("OPTIONS /broths", handlers.HandleBrothOptions)
	router.HandleFunc("GET /proteins", handlers.WithRequiredHeaders(handlers.HandleProteinList))
	router.HandleFunc("OPTIONS /proteins", handlers.HandleProteinOptions)
	router.HandleFunc("POST /orders", handlers.WithRequiredHeaders(handlers.HandleOrderNew))

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
	if response.Header.Get(utilapp.HeaderContentType) != want {
		t.Errorf("response did not have %s of %s, got %v", utilapp.HeaderContentType, want, response.Header)
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

func AssertProteinList(t *testing.T, body io.Reader, want []model.Protein) {
	t.Helper()
	var got []model.Protein
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse response from server %q into protein list: '%v'", body, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertOrderResponse(t *testing.T, body io.Reader, want dto.OrderResponse) {
	t.Helper()
	var got dto.OrderResponse
	err := json.NewDecoder(body).Decode(&got)

	if err != nil {
		t.Fatalf("unable to parse response from server %q into dto.OrderResponse: '%v'", body, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertHeader(t *testing.T, response *http.Response, wantHeader string, wantValue string) {
	t.Helper()
	if response.Header.Get(wantHeader) != wantValue {
		t.Errorf("response did not have %s of %s, got %v", wantHeader, wantValue, response.Header)
	}
}
