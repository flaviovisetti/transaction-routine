package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/entity"

	"github.com/gorilla/mux"
)

// CreateAccount endpoint responsible to receive account and create it
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entity.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	json.NewEncoder(w).Encode(account)
}

// GetAccount endpoint responsible to response and account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}
