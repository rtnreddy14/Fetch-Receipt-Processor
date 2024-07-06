package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"receipt-processor/models"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// receipts is a thread-safe map that stores receipt data along with calculated points.
var receipts sync.Map

// ProcessReceipt handles incoming POST requests to process new receipts.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	// Decode the JSON body into the Receipt struct.
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		log.Println("Error decoding receipt:", err)
		http.Error(w, "Invalid request body", 400)
		return
	}

	// Generate a new UUID for the receipt.
	receiptID := uuid.New()

	// Asynchronously calculate points and store the receipt data.
	go func(r models.Receipt, id uuid.UUID) {
		points := CalcRewardPoints(r) // Perform point calculation in the background.
		receipts.Store(id, models.ReceiptData{Receipt: r, Points: points})
	}(receipt, receiptID)

	// Respond with the generated UUID for the receipt.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: receiptID,
	})
}

// fetchReceiptData retrieves ReceiptData from the receipts map given a string UUID.
func fetchReceiptData(idStr string) (models.ReceiptData, error) {
	var receiptData models.ReceiptData

	id, err := uuid.Parse(idStr)
	if err != nil {
		return receiptData, fmt.Errorf("invalid UUID format")
	}

	value, ok := receipts.Load(id)
	if !ok {
		return receiptData, fmt.Errorf("no receipt found for the given ID")
	}

	receiptData = value.(models.ReceiptData)
	return receiptData, nil
}

// GetPoints returns the points for a receipt given its UUID.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	receiptData, err := fetchReceiptData(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Points uint64 `json:"points"`
	}{
		Points: receiptData.Points,
	})
}

// GetReceiptData returns the complete receipt data for a given receipt ID.
func GetReceiptData(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	receiptData, err := fetchReceiptData(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receiptData)
}
