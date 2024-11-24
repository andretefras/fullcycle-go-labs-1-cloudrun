package presentation

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_InvalidMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestHandler_InvalidBody(t *testing.T) {
	req, err := http.NewRequest("GET", "/", bytes.NewBuffer([]byte("{invalid json}")))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}

func TestHandler_ValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", bytes.NewBuffer([]byte(`{"zipcode":"21930190"}`)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	var weather = &struct {
		City       string  `json:"city"`
		Celsius    float64 `json:"temp_c"`
		Fahrenheit float64 `json:"temp_f"`
		Kelvin     float64 `json:"temp_k"`
	}{}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
	if weather.City == "" {
		t.Error("handler returned empty city")
	}
}
