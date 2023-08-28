package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/broker"
	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"github.com/BaseMax/RabbitMQOrderGo/helpers"
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

	order, _ := models.GetOrderById(uint(orderId))

	_, username := helpers.GetLoggedinUserInfo(c)
	adminName, _, adminMail := conf.GetAdminInfo()
	if username != adminName {
		sub := "User " + username + " created a new refund"
		body := "Hi there admin!\n" +
			"User " + username + " created a refund with following order description:\n" +
			order.Description + "\n"

		go func() {
			err = helpers.EasySendMail(sub, body, adminMail)
			if err != nil {
				log.Println("mailer:", err)
			}
		}()
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

	uid, username := helpers.GetLoggedinUserInfo(c)
	adminName, _, adminMail := conf.GetAdminInfo()

	user, err := models.GetUserById(uid)
	if err != nil {
		return echo.ErrNotFound
	}

	order, err := models.GetOrderById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	sub := "Your refund has has not beed accepted"
	body := "Hi there " + username + "!\n" +
		"Unfortunately your refund has not beed accepted. order description:\n" +
		order.Description + "\n"

	mailAddress := adminMail
	if username == adminName {
		mailAddress = user.Email
	}

	go func() {
		err = helpers.EasySendMail(sub, body, mailAddress)
		if err != nil {
			log.Println("mailer:", err)
		}
	}()

	return c.NoContent(http.StatusNoContent)
}

func handleRefundQueue(c echo.Context, processComplete bool) (*models.Refund, error) {
	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return nil, echo.ErrInternalServerError
		}
	}

	refund, err := broker.DequeueFirstRefund(processComplete)
	if err != nil {
		return nil, echo.ErrNotFound
	}

	if refund == nil {
		return nil, echo.ErrNotFound
	}

	return refund, c.JSON(http.StatusOK, refund)
}

func FirstRefund(c echo.Context) error {
	_, err := handleRefundQueue(c, false)
	return err
}

func CompleteRefund(c echo.Context) error {
	refund, err := handleRefundQueue(c, true)
	if err != nil {
		return err
	}

	sub := "Your refund is complete"
	body := "Hi there " + refund.Order.User.Username + "!\n" +
		"We checked your refund and aggree with it. Order description:\n" +
		refund.Order.Description + "\n" +
		"Please contact us to refund.\n"

	go func() {
		err = helpers.EasySendMail(sub, body, refund.Order.User.Email)
		if err != nil {
			log.Println("mailer:", err)
		}
	}()

	return err
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
