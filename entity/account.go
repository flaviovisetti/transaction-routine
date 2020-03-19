package entity

// Account represents an persist entity.
type Account struct {
	ID             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}
