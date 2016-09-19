package controller

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
)

type JobInfoController struct {
	beego.Controller
}

func (this *JobInfoController)List() {

	infos,_ := service.JobInfoService.List()
	this.Data["infos"] = infos
	this.TplName = "jobinfo/info.html"
}