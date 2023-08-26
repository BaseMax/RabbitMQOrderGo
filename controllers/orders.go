package controllers

import (
	"encoding/json"
	"log"
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
		Status:      models.ORDER_STATUS_PROCESSING,
	}

	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return echo.ErrInternalServerError
		}
	}
	if broker.EnqueueToRabbit(order) != nil {
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

func FetchAllOrders(c echo.Context) error {
	orders, err := models.GetAllOrders()
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, orders)
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
	if models.UpdateOrder(uint(id), order.Description, "") == 0 {
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
	if models.UpdateOrder(uint(id), "", models.ORDER_STATUS_CANCELED) == 0 {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func handleQueue(c echo.Context, processComplete bool) error {
	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return echo.ErrInternalServerError
		}
	}

	order, err := broker.DequeueFirstOrder(processComplete)
	if err != nil {
		log.Println(err)
		return echo.ErrNotFound
	}

	if order == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, order)
}

func FirstOrder(c echo.Context) error {
	return handleQueue(c, false)
}

func CompleteOrder(c echo.Context) error {
	return handleQueue(c, true)
}

func DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if models.DeleteOrder(uint(id)) == 0 {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}
