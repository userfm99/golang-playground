package routers

import (
	"github.com/userfm99/golang-playground/play-beego/controller"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/home",
			beego.NSInclude(
				&controller.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
