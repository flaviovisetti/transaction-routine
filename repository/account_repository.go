package repository

import (
	"github.com/flaviovisetti/transaction-routine/db"
	"github.com/flaviovisetti/transaction-routine/entity"
)

// GetAllAccounts returns a slice of all acounts
func GetAllAccounts() []entity.Account {
	var accounts []entity.Account
	dbase := db.GetConnection()
	defer dbase.Close()

	rows, err := dbase.Query("SELECT * FROM accounts")

	if err != nil {
		panic(err.Error())
	}

	var account entity.Account
	for rows.Next() {
		var id int
		var documentNumber string

		err = rows.Scan(&id, &documentNumber)

		if err != nil {
			panic(err.Error())
		}

		account.ID = id
		account.DocumentNumber = documentNumber

		accounts = append(accounts, account)
	}

	return accounts
}

// GetAccountByID returns an account by ID
func GetAccountByID(accountID int) entity.Account {
	dbase := db.GetConnection()
	defer dbase.Close()

	rows, err := dbase.Query("SELECT * FROM accounts ac WHERE ac.id = $1", accountID)

	if err != nil {
		panic(err.Error())
	}

	var account entity.Account
	for rows.Next() {
		var id int
		var documentNumber string

		err = rows.Scan(&id, &documentNumber)
		account.ID = id
		account.DocumentNumber = documentNumber
	}

	if err != nil {
		panic(err.Error())
	}

	return account
}

// CreateAccount create accounts and return object
func CreateAccount(documentNumber string) entity.Account {
	dbase := db.GetConnection()
	defer dbase.Close()

	var accountID int
	err := dbase.QueryRow("INSERT INTO accounts(document_number) VALUES($1) RETURNING id;", documentNumber).Scan(&accountID)

	if err != nil {
		panic(err.Error())
	}

	account := entity.Account{ID: accountID, DocumentNumber: documentNumber}

	return account
}
