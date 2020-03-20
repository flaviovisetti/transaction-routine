package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flaviovisetti/transaction-routine/entity"

	"github.com/flaviovisetti/transaction-routine/repository"

	"github.com/gorilla/mux"
)

// CreateAccount endpoint responsible to receive account and create it
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entity.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	account = repository.CreateAccount(account.DocumentNumber)
	json.NewEncoder(w).Encode(account)
}

// GetAccount endpoint responsible to response and account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err.Error())
	}

	account := repository.GetAccountByID(accountID)

	json.NewEncoder(w).Encode(account)
}
