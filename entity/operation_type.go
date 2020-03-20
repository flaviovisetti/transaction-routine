package entity

import "github.com/jinzhu/gorm"

// OperationType represents an persist entity.
type OperationType struct {
	gorm.Model
	Description string `gorm:"column:description" json:"description"`
}
