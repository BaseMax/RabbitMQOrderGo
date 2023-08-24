package conf

import "os"

func GetHttpServerAddr() string {
	address := os.Getenv("HTTP_SERVER_ADDRESS")
	if address == "" {
		address = ":8000"
	}
	return address
}
