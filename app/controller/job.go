package controller

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
)

type JobInfoController struct {
	beego.Controller
}

// 查询任务列表
func (this *JobInfoController)List() {

	infos, _ := service.JobInfoService.List()
	this.Data["infos"] = infos
	this.TplName = "jobinfo/list.html"
}

// 跳转到新增任务页面
func (this *JobInfoController)ToAdd() {

	this.TplName = "jobinfo/add.html"

}

func (this *JobInfoController)Add() {

}

type HomeController struct {
	beego.Controller
}

func (this * HomeController) Index() {

	this.TplName = "index.html";
}