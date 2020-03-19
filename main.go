package main

import (
	"log"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
