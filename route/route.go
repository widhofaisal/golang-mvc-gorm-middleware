package route

import (
	"learn/mvc/controller"
	m "learn/mvc/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUser)

	// Logger
	m.LogMiddleware(e)

	// Auth Basic
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUsersController)

	return e
}
