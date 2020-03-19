package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/entity"
)

// CreateTransaction endpoint to create a new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entity.Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	json.NewEncoder(w).Encode(transaction)
}
