package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/v2/home/index", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]string{"message": "this response from v2 port 1881", "time":time.Now().String()})
	})
	e.Start(":1881")
}
