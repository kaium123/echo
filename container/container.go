package container

import (
	"github.com/kaium123/practice/controller"
	"github.com/kaium123/practice/database"
	"github.com/kaium123/practice/repository"
	"github.com/kaium123/practice/router"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	db := database.GetDB()

	productRepository := repository.NewProductsRepository(db)

	productController := controller.NewProductsController(productRepository)

	router.NewProductsRouter(e, productController)

}
