package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/broker"
	"github.com/BaseMax/RabbitMQOrderGo/conf"
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
	if broker.EnqueueToRabbit(order, broker.QUEHE_NAME_ORDERS) != nil {
		return echo.ErrInternalServerError
	}

	_, username := helpers.GetLoggedinUserInfo(c)
	adminName, _, adminMail := conf.GetAdminInfo()
	if username != adminName {
		sub := "User " + username + " created a new order"
		body := "Hi there admin!\n" +
			"User " + username + " created an order with following description:\n" +
			order.Description + "\n"

		go func() {
			err = helpers.EasySendMail(sub, body, adminMail)
			if err != nil {
				log.Println("mailer:", err)
			}
		}()
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

	sub := "Your order has been canceled"
	body := "Hi there " + username + "!\n" +
		"Unfortunately your following order was canceled with following description:\n" +
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

func handleOrderQueue(c echo.Context, processComplete bool) (*models.Order, error) {
	if broker.IsClosed() {
		if broker.ConnectAndCreateQueue() != nil {
			return nil, echo.ErrInternalServerError
		}
	}

	order, err := broker.DequeueFirstOrder(processComplete)
	if err != nil {
		return nil, echo.ErrNotFound
	}

	if order == nil {
		return nil, echo.ErrNotFound
	}

	return order, c.JSON(http.StatusOK, order)
}

func FirstOrder(c echo.Context) error {
	_, err := handleOrderQueue(c, false)
	return err
}

func CompleteOrder(c echo.Context) error {
	order, err := handleOrderQueue(c, true)
	if err != nil {
		return err
	}

	user, err := models.GetUserById(order.UserID)
	if err != nil {
		return echo.ErrNotFound
	}

	sub := "Your order is complete"
	body := "Hi there " + user.Username + "!\n" +
		"We checked your order and aggree with it. Order description:\n" +
		order.Description + "\n"

	go func() {
		err = helpers.EasySendMail(sub, body, user.Email)
		if err != nil {
			log.Println("mailer:", err)
		}
	}()

	return nil
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
