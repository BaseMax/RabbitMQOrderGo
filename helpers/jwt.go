package helpers

import (
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetLoggedinUserInfo(c echo.Context) (uint, string) {
	bearer := c.Request().Header.Get("Authorization")
	token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	username := claims["iss"].(string)
	id, _ := strconv.Atoi(claims["jti"].(string))
	return uint(id), username
}
