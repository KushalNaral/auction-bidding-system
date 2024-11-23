package main

import (
	"log"
	"net/http"
)

func CheckHealth() {
	log.Printf("Starting Bidding Service on port 8080")

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bidding Service is Running...."))
	})

	if err := http.ListenAndServe(":8086", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func main() {

	CheckHealth()
	// Responsible for price generation, validation, and formatting responses
}

func GeneratePrices() {
}

func ValidateBid() {
}

func FormatResponse() {
}
