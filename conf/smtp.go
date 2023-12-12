package conf

import "os"

type MailConf struct {
	FromAddress string
	FromName    string
	Password    string
	Host        string
	Port        string
	Server      string
}

func GetMailConfig() MailConf {
	return MailConf{
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        os.Getenv("MAIL_PORT"),
		Server:      os.Getenv("MAIL_SERVER"),
	}
}
