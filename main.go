package main

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
)

func main() {

	// 初始化服务
	service.Init()
	beego.SetStaticPath("static","assets")
	beego.SetViewsPath("views")
	beego.Run()


}