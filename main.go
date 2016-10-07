package main

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
	"kitty/app/controller"
)

func main() {

	// 初始化服务
	service.Init()
	beego.SetStaticPath("static", "assets")
	beego.SetViewsPath("views")
	beego.Router("/", &controller.HomeController{}, "*:Index")
	beego.Router("/jobinfo/list", &controller.JobInfoController{}, "*:List")
	beego.Router("/jobinfo/add", &controller.JobInfoController{}, "get:ToAdd")
	beego.Router("/jobinfo/add", &controller.JobInfoController{}, "post:Add")
	beego.Router("/jobinfo/edit", &controller.JobInfoController{}, "*:Edit")
	beego.Router("/jobinfo/info",&controller.JobInfoController{},"*:Info")
	beego.Router("/jobinfo/delete",&controller.JobInfoController{},"*:Delete")

	beego.Router("/jobsanpshot/list",&controller.JobSanpshotController{},"*:List")


	beego.Run()

}