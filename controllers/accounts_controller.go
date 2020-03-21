package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/db"
	"github.com/flaviovisetti/transaction-routine/entity"
	"github.com/gorilla/mux"
)

// CreateAccount endpoint responsible to receive account and create
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var account entity.Account
	json.NewDecoder(r.Body).Decode(&account)
	db.DBCon.Create(&account)

	json.NewEncoder(w).Encode(account)
}

// GetAccount endpoint responsible to response and account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var account entity.Account
	db.DBCon.First(&account, params["id"])

	json.NewEncoder(w).Encode(account)
}
