package routes

import (
	"crud_101/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:productID", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
}
