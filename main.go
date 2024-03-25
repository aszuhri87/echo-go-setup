package main

import (
	"echo-go/config"
	_ "echo-go/docs"
	"echo-go/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:1323
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description use `Bearer <xx token xx>` to authenticate
func main() {
	config.Conn()
	config.InitDB()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.AuthRoutes(e)
	routes.UserRoutes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// func handleArgs() {
// 	flag.Parse()
// 	args := flag.Args()

// 	// user := User{Name: "admin"}
// 	// config.DB.First(&user)

// 	// if len(user) >= 1 {
// 	// 	os.Exit(0)
// 	// }

// 	if len(args) >= 1 {
// 		switch args[0] {
// 		case "seed":
// 			// connect DB
// 			seeds.Execute(config.DB, args[1:]...)
// 			os.Exit(0)
// 		}
// 	}
// }
