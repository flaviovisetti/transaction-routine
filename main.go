package main

import (
	"log"
	"net/http"

	"github.com/flaviovisetti/transaction-routine/controllers"
	"github.com/flaviovisetti/transaction-routine/db"
	"github.com/flaviovisetti/transaction-routine/migrations"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	migrations.Migrate()
	defer db.DBCon.Close()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
