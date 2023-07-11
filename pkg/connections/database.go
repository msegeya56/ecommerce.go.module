package connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	
)



var Db *gorm.DB

// NewSQLiteConnection creates a new SQLite database connection using go


func NewSQLiteConnection(dbPath string) (*gorm.DB, error) {
Db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate all the eCommerce structs
	err = Db.AutoMigrate(
		&entities.Category{},
		&entities.Checkout{},
		&entities.CreditLimit{},
		&entities.Customer{},
		&entities.Deposit{},
		&entities.Invoice{},
		&entities.Order{},
		&entities.Payment{},
		&entities.Product{},
		&entities.Receipt{},
		&entities.Review{},
		&entities.Rating{},
		&entities.CreditCard{},
	)
	if err != nil {
		return nil, err
	}

	return Db, nil
}