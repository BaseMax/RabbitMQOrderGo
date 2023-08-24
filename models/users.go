package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Role string `gorm:"not null"`
}
