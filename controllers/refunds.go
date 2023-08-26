package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/broker"
	"github.com/BaseMax/RabbitMQOrderGo/models"
)

func CreateRefund(c echo.Context) error {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	refunId, err := models.CreateRefund(uint(orderId))
	if err != nil {
		if err.Error() == models.REFUND_ERROR_BADORDER {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": models.REFUND_ERROR_BADORDER,
			})
		}
		return echo.ErrConflict
	}

	refund := models.Refund{
		ID:      refunId,
		OrderID: uint(orderId),
		Status:  models.REFUND_STATUS_APPENDING,
	}

	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return echo.ErrInternalServerError
		}
	}
	if broker.EnqueueToRabbit(refund, broker.QUEHE_NAME_REFUNDS) != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, refund)
}

func FetchRefund(c echo.Context) error {
	return nil
}

func FetchAllRefunds(c echo.Context) error {
	return nil
}

func RefundStatus(c echo.Context) error {
	return nil
}

func CancelRefund(c echo.Context) error {
	return nil
}

func FirstRefund(c echo.Context) error {
	return nil
}

func CompleteRefund(c echo.Context) error {
	return nil
}

func DeleteRefund(c echo.Context) error {
	return nil
}
