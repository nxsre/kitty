package controller

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController)WriteJson(v interface{}) {

	this.Data["json"] = v
	this.ServeJSON(false, false)
	this.StopRun()

}