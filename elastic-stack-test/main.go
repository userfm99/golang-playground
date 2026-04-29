package main

import (
	"github.com/kataras/iris"
	"github.com/userfm99/golang-playground/elastic-stack-test/logx"
)

type Req struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	app := iris.Default()
	app.Get("/elk-log-test", home)
	app.Post("/add", add)
	if err := app.Run(iris.Addr(":8832")); err != nil {
		panic(err)
	}
}

func home(ctx iris.Context) {
	logx.Info("Request to home")
	ctx.StatusCode(200)
	ctx.JSON(iris.Map{
		"greeting": "Welcome",
	})
}

func add(ctx iris.Context) {
	var req Req
	if err := ctx.ReadJSON(&req); err != nil {
		logx.Err(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": err.Error()})
		return
	}
	logx.Debug("Request to add from ip " + ctx.Request().RemoteAddr)
	ctx.StatusCode(200)
	ctx.JSON(iris.Map{"message": "request accepted"})
}
