package commons

import (
	"time"

	"gorm.io/gorm"
)

type Foundation struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Type        string
	Stage       string
	Status      int
	Data        []byte
	Document    []byte
	Events      []byte
	CreateAt    time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Timestamp   time.Time
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
	Timestamp   time.Time
}
