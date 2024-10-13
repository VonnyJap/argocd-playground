package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemperatureHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/temperature", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(temperatureHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Temperature{
		City:       "San Francisco",
		Temperature: 68.0,
		Unit:       "Fahrenheit",
	}

	var response Temperature
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if response != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", response, expected)
	}
}
