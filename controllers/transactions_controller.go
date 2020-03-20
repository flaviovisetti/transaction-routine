package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/db"
	"github.com/flaviovisetti/transaction-routine/entity"
)

// CreateTransaction endpoint to create a new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transaction entity.Transaction
	json.NewDecoder(r.Body).Decode(&transaction)
	db.DBCon.Create(&transaction)

	json.NewEncoder(w).Encode(transaction)
}
