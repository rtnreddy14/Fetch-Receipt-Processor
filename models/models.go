package models

// Receipt defines the structure for storing receipt information.
type Receipt struct {
	Retailer     string  `json:"retailer"`     // Name of the retailer.
	PurchaseDate string  `json:"purchaseDate"` // Date of the purchase.
	PurchaseTime string  `json:"purchaseTime"` // Time of the purchase.
	Items        []Item  `json:"items"`        // List of items purchased.
	Total        float64 `json:"total,string"` // Total amount of the receipt, expected as a string in JSON.
}

// Item represents a single item on a receipt.
type Item struct {
	ShortDescription string  `json:"shortDescription"` // Brief description of the item.
	Price            float64 `json:"price,string"`     // Price of the item, expected as a string in JSON.
}

// ReceiptData combines a Receipt with its calculated reward points.
type ReceiptData struct {
	Receipt Receipt `json:"receipt"` // The receipt details.
	Points  uint64  `json:"points"`  // Calculated points based on the receipt.
}
