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
