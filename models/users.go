package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Password string `gorm:"not null" json:"pass"`
	Username string `gorm:"unique;not null" json:"user"`
	Email    string `gorm:"unique;not null" json:"email"`
}

func RegisterUser(username, password, email string) error {
	hashedByte := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedByte[:])
	user := User{Username: username, Password: hashedString, Email: email}
	return db.Create(&user).Error
}

func LoginUser(username, password, email string) bool {
	hashedByte := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedByte[:])

	var count int64
	db.Model(&User{}).Where("username = ? AND password = ? AND email = ?", username, hashedString, email).Count(&count)
	return count == 1
}
