package handlers

import (
	"math"
	"receipt-processor/models"
	"strings"
	"time"
)

// CalcRewardPoints calculates and returns the reward points for a given receipt.
func CalcRewardPoints(receipt models.Receipt) uint64 {
	var points uint64

	// Add points for each alphanumeric character in the retailer's name.
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points += 1
		}
	}

	// Add 50 points if the total amount is a whole number (no cents).
	if receipt.Total == float64(int64(receipt.Total)) {
		points += 50
	}

	// Add 25 points if the total is a multiple of 0.25.
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// Add 5 points for every two items in the receipt.
	points += uint64((len(receipt.Items) / 2) * 5)

	// Calculate additional points based on item descriptions and prices.
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if math.Mod(float64(len(trimmedDescription)), 3) == 0 {
			points += uint64(math.Ceil(item.Price * 0.2))
		}
	}

	// Add points based on the date and time specifics of the purchase.
	receiptDateTime, _ := time.Parse("2006-01-02 15:04", receipt.PurchaseDate+" "+receipt.PurchaseTime)
	if math.Mod(float64(receiptDateTime.Day()), 2) == 1 {
		points += 6
	}
	if receiptDateTime.Hour() >= 14 && receiptDateTime.Hour() < 16 {
		points += 10
	}

	return points
}
