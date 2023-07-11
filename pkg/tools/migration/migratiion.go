package migration

import (
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/entities"
	"gorm.io/gorm"
)




func SyncToDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.Customer{},
		&entities.Category{},
		&entities.Checkout{},
		&entities.CreditLimit{},
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
		return err
	}

	return nil
}





