package entity

import "github.com/jinzhu/gorm"

// Account represents an persist entity.
type Account struct {
	gorm.Model
	DocumentNumber string `gorm:"column:document_number" json:"document_number"`
}
