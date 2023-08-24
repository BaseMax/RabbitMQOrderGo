package conf

import (
	"fmt"
	"os"
)

func GetRabbitUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:5672",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOSTNAME"),
	)
}
