package migrations

import (
	"github.com/flaviovisetti/transaction-routine/db"
	"github.com/flaviovisetti/transaction-routine/entity"
)

// Migrate execute a gorm AutoMigrate
func Migrate() {
	db.DBCon.AutoMigrate(
		entity.Account{},
		entity.OperationType{},
		entity.Transaction{},
	)

	db.DBCon.Model(&entity.Transaction{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "RESTRICT")
	db.DBCon.Model(&entity.Transaction{}).AddForeignKey("operation_type_id", "operation_types(id)", "CASCADE", "RESTRICT")
}
