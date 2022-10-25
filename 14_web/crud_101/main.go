package main

import (
	"crud_101/configs"
	"crud_101/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Println("Hello world")

	// connect to DB
	configs.ConnectDB()

	// routes
	routes.UserRoutes(e)
	routes.ProductRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Echo and MongoDB"})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
