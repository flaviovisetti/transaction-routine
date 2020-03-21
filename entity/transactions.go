package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction represents an persist entity.
type Transaction struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time  `json:"-"`
	UpdatedAt       time.Time  `json:"-"`
	DeletedAt       *time.Time `json:"-"`
	AccountID       int        `gorm:"column:account_id" json:"account_id"`
	OperationTypeID int        `gorm:"column:operation_type_id" json:"operation_type_id"`
	Amount          float64    `gorm:"column:amount" json:"amount"`
	EventDate       time.Time  `gorm:"column:event_date" json:"event_date"`
}

// BeforeCreate validates amount and operation_type
func (t *Transaction) BeforeCreate(scope *gorm.Scope) error {
	if t.OperationTypeID != 4 {
		scope.SetColumn("Amount", t.Amount*-1)
	}

	scope.SetColumn("EventDate", time.Now())
	return nil
}
