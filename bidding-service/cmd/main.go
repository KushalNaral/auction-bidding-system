package main

import (
	"log"
	"net/http"
)

func main() {
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
