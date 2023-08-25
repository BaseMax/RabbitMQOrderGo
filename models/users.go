package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Password string `gorm:"not null"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

func RegisterUser(username, password, email string) error {
	hashedByte := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedByte[:])
	user := User{Username: username, Password: hashedString, Email: email}
	return db.Create(&user).Error
}
