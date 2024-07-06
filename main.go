package main

import (
	"log"
	"net/http"
	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register a route for processing receipts via POST requests.
	r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")

	// Register a route to retrieve points by receipt ID via GET requests.
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	// Register a route to fetch complete receipt data by receipt ID via GET requests.
	r.HandleFunc("/receipts/{id}", handlers.GetReceiptData).Methods("GET")

	// Start the HTTP server on port 8080 and log if there is an error.
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
