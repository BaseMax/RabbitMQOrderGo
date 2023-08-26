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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	refund, err := models.GetRefundById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, refund)
}

func FetchAllRefunds(c echo.Context) error {
	refunds, err := models.GetAllRefunds()
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, refunds)
}

func RefundStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	refund, err := models.GetRefundById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]any{"status": refund.Status})
}

func CancelRefund(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if models.UpdateRefund(uint(id), models.REFUND_STATUS_CANCELED) == 0 {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func handleRefundQueue(c echo.Context, processComplete bool) error {
	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return echo.ErrInternalServerError
		}
	}

	refund, err := broker.DequeueFirstRefund(processComplete)
	if err != nil {
		return echo.ErrNotFound
	}

	if refund == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, refund)
}

func FirstRefund(c echo.Context) error {
	return handleRefundQueue(c, false)
}

func CompleteRefund(c echo.Context) error {
	return handleRefundQueue(c, true)
}

func DeleteRefund(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if models.DeleteRefund(uint(id)) == 0 {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}
