package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Temperature struct {
	City       string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Unit       string  `json:"unit"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	tempData := Temperature{
		City:       "San Francisco",
		Temperature: 68.0,
		Unit:       "Fahrenheit",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tempData)
}

func main() {
	http.HandleFunc("/temperature", temperatureHandler)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
