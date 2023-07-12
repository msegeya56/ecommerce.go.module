package migration

import (
	"github.com/msegeya56/ecommerce.go.module/pkg/domains/models"

)




func SyncToDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Customer{},
		&models.Category{},
		&models.Checkout{},
		&models.CreditLimit{},
		&models.Deposit{},
		&models.Invoice{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
		&models.Receipt{},
		&models.Review{},
		&models.Rating{},
		&models.CreditCard{},
	)
	if err != nil {
		return err
	}

	return nil
}





