package models

type Product struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint
	Name    string  `gorm:"not null"`
	Price   float64 `gorm:"not null"`
}
