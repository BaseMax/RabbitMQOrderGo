package models

type Refund struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"unique"`
	Order   Order
	Status  string `gorm:"not null"`
}
