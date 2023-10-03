package commons

import (
	"database/sql"

	"gorm.io/gorm"
)

type DeletedAt sql.NullTime

type Foundation struct {
	gorm.Model
	Name        string
	Description string
	Type        string
	Stage       string
	Status      int
	Data        []byte
	Document    []byte
	Events      []byte
	DeletedAt   DeletedAt `gorm:"index"`
}

type FoundationEntity struct {
	gorm.Model
	Name        string
	Description string
	Type        string
	Stage       string
	Status      int
	Data        []byte
	Document    []byte
	Events      []byte
	DeletedAt   DeletedAt `gorm:"index"`
}
