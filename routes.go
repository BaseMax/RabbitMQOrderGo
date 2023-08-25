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

	g.POST("refunds/", TODO)
	g.GET("refunds/:id", TODO)
	g.GET("refunds/:id/status", TODO)
	g.POST("refunds/:id/decancel", TODO)
	g.GET("refunds/last", TODO, middlewares.IsAdmin)
	g.POST("refunds/last/done", TODO, middlewares.IsAdmin)
	g.DELETE("refunds/:id", TODO, middlewares.IsAdmin)

	return e
}
