package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting Auction Service on port 8081")

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Auctions Service is Running...."))
	})

	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
