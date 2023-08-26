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

func LoginUser(username, password, email string) (uint, error) {
	hashedByte := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedByte[:])

	var user User
	err := db.First(&user).Where("username = ? AND password = ? AND email = ?", username, hashedString, email).Error
	return user.ID, err
}

func GetAllUsers() (users []User, err error) {
	err = db.Select("id, username, email").Find(&users).Error
	return
}

func UpdateUser(id uint, username, password, email string) int64 {
	hashedByte := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedByte[:])
	return db.Where(id).Updates(&User{Username: username, Password: hashedString, Email: email}).RowsAffected
}
