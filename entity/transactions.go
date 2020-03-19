package entity

// Transaction represents an persist entity.
type Transaction struct {
	ID            int
	Account       Account
	OperationType OperationType
	Amount        float64
	EventDate     string
}
