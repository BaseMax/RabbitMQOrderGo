package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"github.com/BaseMax/RabbitMQOrderGo/helpers"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		_, loggedinName := helpers.GetLoggedinUserInfo(c)
		name, _, _ := conf.GetAdminInfo()
		if loggedinName != name {
			return echo.ErrUnauthorized
		}
		return next(c)
	})
}
