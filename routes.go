package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/RabbitMQOrderGo/conf"
	"github.com/BaseMax/RabbitMQOrderGo/controllers"
	"github.com/BaseMax/RabbitMQOrderGo/middlewares"
)

func TODO(c echo.Context) error { return nil }

func initRoutes() *echo.Echo {
	e := echo.New()
	g := e.Group("/", echojwt.WithConfig(echojwt.Config{SigningKey: conf.GetJwtSecret()}))

	g.GET("health", controllers.HealthCheck, middlewares.IsAdmin)

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	g.POST("refresh", controllers.Refresh)
	g.PUT("user/:username", TODO, middlewares.IsAdmin)

	g.POST("orders", controllers.CreateOrder)
	g.GET("orders/:id", controllers.FetchOrder)
	g.GET("orders", controllers.FetchAllOrders)
	g.PUT("orders/:id", controllers.EditOrder)
	g.GET("orders/:id/status", controllers.OrderStatus)
	g.POST("orders/:id/cancel", controllers.CancelOrder)
	g.GET("orders/fist", controllers.FirstOrder, middlewares.IsAdmin)
	g.POST("orders/first/done", controllers.CompleteOrder, middlewares.IsAdmin)
	g.DELETE("orders/:id", controllers.DeleteOrder, middlewares.IsAdmin)

	g.POST("refunds", controllers.CreateRefund)
	g.GET("refunds/:id", controllers.FetchRefund)
	g.GET("refunds/", controllers.FetchAllRefunds)
	g.GET("refunds/:id/status", controllers.RefundStatus)
	g.POST("refunds/:id/cancel", controllers.CancelRefund)
	g.GET("refunds/first", controllers.FirstRefund, middlewares.IsAdmin)
	g.POST("refunds/last/done", controllers.CompleteRefund, middlewares.IsAdmin)
	g.DELETE("refunds/:id", controllers.DeleteRefund, middlewares.IsAdmin)

	return e
}
