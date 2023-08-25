package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		if bearer == "" {
			return echo.ErrBadRequest
		}
		token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
		claims := token.Claims.(jwt.MapClaims)

		name, _, _ := conf.GetAdminInfo()
		if claims["iss"] != name {
			return echo.ErrUnauthorized
		}
		return next(c)
	})
}
