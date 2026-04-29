package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customeErrorHandler

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/home", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"Say": "hello"})
	})

	e.Server.Addr = ":8786"

	graceful.ListenAndServe(e.Server, 10*time.Second)
}

func customeErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal server error"
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		msg = httpErr.Message.(string)
	}
	c.JSON(code, echo.Map{"error": msg})
}
