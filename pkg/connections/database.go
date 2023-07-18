package connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/msegeya56/ecommerce.go.module/pkg/domains/models"
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
		&models.Category{},
		&models.Checkout{},
		&models.Creditlimit{},
		&models.Customer{},
		&models.Deposit{},
		&models.Invoice{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
		&models.Receipt{},
		&models.Review{},
		&models.Rating{},
		&models.Creditcard{},
	)
	if err != nil {
		return nil, err
	}

	return Db, nil
}

