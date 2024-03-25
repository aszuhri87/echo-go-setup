package routes

import (
	"echo-go/app/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) *echo.Echo {
	r := e.Group("/user")

	// middleware
	// r.Use(config.JwtMiddlewareSet())

	r.GET("", controllers.GetUser)
	r.POST("", controllers.CreateUser)
	r.GET("/:id", controllers.GetUserByID)
	r.GET("/profile", controllers.UserProfile)
	r.PUT("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)

	return e
}
