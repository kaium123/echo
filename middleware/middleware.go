package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Auth(c *echo.Echo) echo.MiddlewareFunc {
	f := middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "jack" && password == "123" {
			return true, nil
		}

		return false, nil
	})
	return f
}
