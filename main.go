package main

import (
	"github.com/astaxie/beego"
	"kitty/app/action"
)

func main() {

	beego.SetStaticPath("static","assets")
	beego.SetViewsPath("views")
	beego.Router("/",&action.IndexController{},"*:Index")
	beego.Run()


}