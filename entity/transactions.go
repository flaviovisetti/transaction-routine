package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction represents an persist entity.
type Transaction struct {
	gorm.Model
	AccountID       int       `gorm:"column:account_id" json:"account_id"`
	OperationTypeID int       `gorm:"column:operation_type_id" json:"operation_type_id"`
	Amount          float64   `gorm:"column:amount" json:"amount"`
	EventDate       time.Time `gorm:"column:event_date" json:"event_date"`
}
