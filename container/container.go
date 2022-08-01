package container

import (
	"github.com/kaium123/practice/controller"
	"github.com/kaium123/practice/database"
	"github.com/kaium123/practice/repository"
	"github.com/kaium123/practice/router"
	"github.com/kaium123/practice/utility"

	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	db := database.GetDB()

	productRepository := repository.NewProductsRepository(db)

	productUtility := utility.NewProductsUtility(productRepository)

	productController := controller.NewProductsController(productRepository, productUtility)

	router.NewProductsRouter(e, productController)

}
