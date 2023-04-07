package route

import (
	"learn/mvc/constant"
	"learn/mvc/controller"
	m "learn/mvc/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUser)
	e.POST("/login", controller.LoginUserController)

	// Logger
	m.LogMiddleware(e)

	// Auth Basic
	eAuthBasic := e.Group("/auth")							// grouping just for example using middleware
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUsersController)
	
	// Auth JWT
	eJwt := e.Group("/jwt")									// grouping just for example using middleware
	eJwt.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUsersController)

	return e
}
