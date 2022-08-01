package cmd

import (
	"github.com/kaium123/practice/config"
	"github.com/kaium123/practice/container"
	"github.com/kaium123/practice/database"
	"github.com/labstack/echo/v4"
)

func Execute() {
	e := echo.New()
	config.LoadConfig()
	database.InitDB()
	container.Init(e)
	port := config.App().Port
	e.Logger.Fatal(e.Start(":" + port))
}
