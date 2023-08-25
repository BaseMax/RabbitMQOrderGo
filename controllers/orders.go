package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/broker"
	"github.com/BaseMax/RabbitMQOrderGo/helpers"
	"github.com/BaseMax/RabbitMQOrderGo/models"
)

func CreateOrder(c echo.Context) error {
	var order models.Order
	if err := json.NewDecoder(c.Request().Body).Decode(&order); err != nil {
		return echo.ErrBadRequest
	}

	iid, _ := helpers.GetLoggedinUserInfo(c)
	orderId, err := models.CreateOrder(iid, order.Description)
	if err != nil {
		return echo.ErrBadRequest
	}

	order = models.Order{
		ID:          orderId,
		UserID:      iid,
		Description: order.Description,
		Status:      "processing",
	}

	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return echo.ErrInternalServerError
		}
	}
	if broker.EnqueueOrderToRabbit(order) != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, order)
}

func FetchOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	order, err := models.GetOrderById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, order)
}

func EditOrder(c echo.Context) error {
	var order models.Order
	if err := json.NewDecoder(c.Request().Body).Decode(&order); err != nil {
		return echo.ErrBadRequest
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if models.UpdateOrder(uint(id), order.Description, "") != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func OrderStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	order, err := models.GetOrderById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]any{"status": order.Status})
}

func CancelOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if models.UpdateOrder(uint(id), "", "canceled") != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}
