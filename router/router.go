package router

import (

	//"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/controller"
	"github.com/kaium123/practice/database"
	"github.com/labstack/echo/v4"
)

func Route() {

	e := echo.New()
	//e.Validator = &ProductValidator{validator: v}
	e.GET("product/:id", controller.Get)
	e.POST("product", controller.Post)
	e.PUT("product/:id", controller.Update)
	e.DELETE("product/:id", controller.Delete)
	e.GET("product", controller.Getall)
	e.Logger.Fatal(e.Start(":8000"))
	database.InitDB()
}
