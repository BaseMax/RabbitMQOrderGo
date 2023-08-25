package models

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Status      string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`

	User User
}
