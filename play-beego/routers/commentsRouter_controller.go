package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/userfm99/golang-playground/play-beego/controller:MainController"] = append(beego.GlobalControllerRouter["github.com/userfm99/golang-playground/play-beego/controller:MainController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/creatures`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/userfm99/golang-playground/play-beego/controller:MainController"] = append(beego.GlobalControllerRouter["github.com/userfm99/golang-playground/play-beego/controller:MainController"],
		beego.ControllerComments{
			Method:           "GetById",
			Router:           `/creatures/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("id", param.InPath),
			),
			Params: nil})

}
