package controller

import (
	"log"

	"github.com/userfm99/golang-playground/play-beego/model"

	"github.com/astaxie/beego"
)

var allCreatures = []model.Creature{
	{
		Id:   1,
		Name: "C1",
	},
	{
		Id:   2,
		Name: "C2",
	},
}

type MainController struct {
	beego.Controller
}

// @Title Get
// @Description get home v5
// @Success 200 string success
// @Failure 403 failed to get data
// @router /creatures [get]
func (this *MainController) GetAll() {
	this.Data["json"] = allCreatures
	this.ServeJSON()
}

// @Title Get
// @Description get home v5
// @Success 200 string success
// @Failure 403 failed to get data
// @router /creatures/:id [get]
func (this *MainController) GetById(id int) {
	log.Println("id request", id)
	this.Data["json"] = allCreatures[id]
	this.ServeJSON()
}
