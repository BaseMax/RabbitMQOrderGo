package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"github.com/BaseMax/RabbitMQOrderGo/models"
)

var EXPTIME = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

func Register(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	if models.RegisterUser(user.Username, user.Password, user.Email) != nil {
		return echo.ErrConflict
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		Issuer:    user.Username,
		ExpiresAt: EXPTIME,
	})
	bearer, _ := token.SignedString([]byte(conf.GetJwtSecret()))
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func Login(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	if !models.LoginUser(user.Username, user.Password, user.Email) {
		return echo.ErrNotFound
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		Issuer:    user.Username,
		ExpiresAt: EXPTIME,
	})
	bearer, _ := token.SignedString([]byte(conf.GetJwtSecret()))
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func Refresh(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	if bearer == "" {
		return echo.ErrBadRequest
	}
	token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(claims["jti"]),
		Issuer:    fmt.Sprint(claims["iss"]),
		ExpiresAt: EXPTIME,
	})
	bearer, _ = refreshToken.SignedString([]byte(conf.GetJwtSecret()))
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}
