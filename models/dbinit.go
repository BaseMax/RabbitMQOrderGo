package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
)

var db *gorm.DB

func InitDB() (err error) {
	dsn := conf.GetMysqlDsn()
	db, err = gorm.Open(mysql.Open(dsn))
	return err
}

func Migrate() error {
	var count int64

	if err := db.AutoMigrate(&User{}, &Order{}, &Refund{}); err != nil {
		return err
	}

	user, pass, email := conf.GetAdminInfo()
	db.Model(&User{}).Where("username = ?", user).Count(&count)
	if count == 1 {
		return nil
	}

	return RegisterUser(user, pass, email)
}
