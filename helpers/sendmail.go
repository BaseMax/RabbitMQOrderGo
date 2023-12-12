package helpers

import (
	"fmt"
	"net/smtp"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
)

var auth smtp.Auth

func EasySendMail(sub, body, to string) error {
	c := conf.GetMailConfig()

	if auth == nil {
		auth = smtp.PlainAuth("", c.FromAddress, c.Password, c.Host)
	}

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+body,
		c.FromAddress, to, sub,
	)

	return smtp.SendMail(c.Server, auth, c.FromAddress, []string{to}, []byte(msg))
}
