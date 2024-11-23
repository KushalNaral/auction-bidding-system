package main

import (
	"log"
	"net/http"
)

func CheckHealth() {
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

func main() {

	CheckHealth()
	// Orchestrates the auction, managing bid collection, aggregating bids, and selecting the winner.
}

func ManageAuction() {
}

func ManageBidCollection() {
}

func AggregateBids() {

}

func SelectWinner() {
}

func ValidateResponse() {
}
