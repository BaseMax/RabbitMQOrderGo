package models

import (
	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (err error) {
	dsn := conf.GetMysqlDsn()
	db, err = gorm.Open(mysql.Open(dsn))
	return err
}
