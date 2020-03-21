package entity

import (
	"time"
)

// Account represents an persist entity.
type Account struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
	DocumentNumber string     `gorm:"column:document_number" json:"document_number"`
}
