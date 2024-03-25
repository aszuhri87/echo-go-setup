package routes

import (
	"echo-go/app/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) *echo.Echo {
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.RegisterUser)

	return e
}
