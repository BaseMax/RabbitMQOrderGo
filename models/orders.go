package models

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Status     string `gorm:"not null"`

	Customer Customer
	Products []Product
}
