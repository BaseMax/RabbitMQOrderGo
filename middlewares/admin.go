package middlewares

import "github.com/labstack/echo/v4"

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		return next(c)
	})
}
