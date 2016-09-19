package main

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
	"kitty/app/controller"
)

func main() {

	// 初始化服务
	service.Init()
	beego.SetStaticPath("static","assets")
	beego.SetViewsPath("views")
	beego.Router("/jobinfo/list",&controller.JobInfoController{},"*:List")
	beego.Router("/jobinfo/add",&controller.JobInfoController{},"get:ToAdd")
	beego.Run()


}