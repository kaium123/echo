package cmd

import (
	"github.com/kaium123/practice/config"
	"github.com/kaium123/practice/container"
	"github.com/kaium123/practice/database"
	"github.com/kaium123/practice/middleware"
	"github.com/labstack/echo/v4"
)

func Execute() {
	e := echo.New()
	if err := middleware.Attach(e); err != nil {
		panic(err)
	}
	config.LoadConfig()
	database.InitDB()
	container.Init(e)
	port := config.App().Port
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
