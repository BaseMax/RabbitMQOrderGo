package conf

import (
	"github.com/joho/godotenv"
)

func Init() error {
	return godotenv.Load(".env")
}
