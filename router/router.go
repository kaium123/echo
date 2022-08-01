package router

import (

	//"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/controller"
	"github.com/labstack/echo/v4"
)

type products struct {
	productC controller.IProducts
}

func NewProductsRouter(e *echo.Echo, productController controller.IProducts) {
	prod := &products{
		productC: productController,
	}

	e.GET("product/:id", prod.productC.GetProduct)
	e.POST("product", prod.productC.Post)
	e.PUT("product/:id", prod.productC.Update)
	e.DELETE("product/:id", prod.productC.Delete)
	e.GET("product", prod.productC.Get)
	//e.GET("product/name", prod.productC.GetProductName)
	//e.GET("product/search", prod.productC.SearchProductName)
}
