package main

import (
	"github.com/astaxie/beego"
	"github.com/soopsio/kitty/app/controller"
	"github.com/soopsio/kitty/app/job"
	"github.com/soopsio/kitty/app/service"
)

func main() {

	// init all service
	service.Init()
	// init quartz module
	//job.BootstrapJobManager()
	job.NewJobManager()
	job.JobManager.PushAllJob()

	beego.SetStaticPath("static", "assets")
	beego.SetViewsPath("views")
	beego.Router("/", &controller.HomeController{}, "*:Index")
	beego.Router("/jobinfo/list", &controller.JobInfoController{}, "*:List")
	beego.Router("/jobinfo/add", &controller.JobInfoController{}, "get:ToAdd")
	beego.Router("/jobinfo/add", &controller.JobInfoController{}, "post:Add")
	beego.Router("/jobinfo/edit", &controller.JobInfoController{}, "*:Edit")
	beego.Router("/jobinfo/info", &controller.JobInfoController{}, "*:Info")
	beego.Router("/jobinfo/delete", &controller.JobInfoController{}, "*:Delete")

	beego.Router("/jobinfo/active", &controller.JobInfoController{}, "*:Active")

	beego.Router("/jobsnapshot/list", &controller.JobSanpshotController{}, "*:List")
	beego.Router("/jobsnapshot/info", &controller.JobSanpshotController{}, "*:Info")

	beego.Router("/jobsnapshot/delete", &controller.JobSanpshotController{}, "*:Delete")

	beego.Router("/monitor/", &controller.MonitorController{}, "*:List")

	beego.AddFuncMap("inc", inc)
	beego.AddFuncMap("sub", sub)
	beego.Run()

}

// 相加
func inc(num, num2 int) (result int) {

	return num + num2
}

// 相减
func sub(num, num2 int) (result int) {

	return num - num2
}
