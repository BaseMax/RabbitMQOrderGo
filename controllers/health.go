package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/broker"
)

func HealthCheck(c echo.Context) error {
	status := broker.GetStatus()
	return c.JSON(http.StatusOK, map[string]any{
		"rabbitmq": status,
	})
}
