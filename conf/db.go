package conf

import (
	"fmt"
	"os"
)

func GetMysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOSTNAME"),
		os.Getenv("MYSQL_DATABASE"),
	)
}

func GetAdminInfo() (user, pass, email string) {
	user = os.Getenv("ADMIN_NAME")
	pass = os.Getenv("ADMIN_PASSWORD")
	email = os.Getenv("ADMIN_EMAIL")
	return
}
