package router

import (
	"github.com/kaium123/practice/controller"
	"github.com/kaium123/practice/middleware"
	"github.com/labstack/echo/v4"
)

func NewProductsRouter(e *echo.Echo, productController controller.IProductsController) {

	e.POST("/products", productController.Create, middleware.Auth(e))
	e.PUT("/products/:id", productController.Update, middleware.Auth(e))
	e.DELETE("/products/:id", productController.Delete, middleware.Auth(e))
	e.GET("products", productController.GetProducts)
	e.GET("products/:id", productController.GetProduct)
}
