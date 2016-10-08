package controller

import (
	"kitty/app/service"
	"github.com/astaxie/beego/validation"
	"kitty/app/common"
)

type JobInfoController struct {
	BaseController
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

// 新增任务job
func (this *JobInfoController)Add() {
	result := common.Result{}

	//
	valid := validation.Validation{}
	//
	jobName := this.GetString("JobName")
	//
	jobGroup := this.GetString("JobGroup")
	//
	cron := this.GetString("Cron")
	//
	url := this.GetString("Url")
	//
	params := this.GetString("Params")
	//
	active, _ := this.GetInt("Active", 0)
	phone := this.GetString("Phone")
	remark := this.GetString("Remark")

	valid.Required(jobName, jobName).Message("任务名称不能为空!")
	valid.Required(jobGroup, jobName).Message("任务名称不能为空!")
	valid.Required(cron, cron).Message("Cron表达式不能为空!")
	valid.Required(url, url).Message("任务服务URL地址不能为空!")
	valid.Required(params, params).Message("任务服务地址参数不能为空!")

	if valid.HasErrors() {

		err := valid.Errors[0]
		result.Message = err.String()

		this.WriteJson(result)

	}

	err := service.JobInfoService.Add(jobName, jobGroup, cron, url, params, phone, remark, active)

	if err != nil {
		result.Message = err.Error()
		this.WriteJson(result)
	} else {
		result.Success = true
		result.Message = "保存成功"
		this.WriteJson(result)
	}
}



// 修改
func (this *JobInfoController)Edit() {

	//
	if this.Ctx.Request.Method == "POST" {
		result := &common.Result{}
		url := this.GetString("Url")
		cron := this.GetString("Cron")
		params := this.GetString("Params")
		remark := this.GetString("Remark")
		phone := this.GetString("Phone")
		id, _ := this.GetInt("Id")

		valid := validation.Validation{}
		valid.Required(cron, cron).Message("Cron表达式不能为空!")
		valid.Required(url, url).Message("任务服务URL地址不能为空!")
		valid.Required(params, params).Message("任务服务地址参数不能为空!")
		if valid.HasErrors() {

			err := valid.Errors[0]
			result.Message = err.String()

			this.WriteJson(result)

		} else {

			err := service.JobInfoService.UpdateJobInfo(id, url, cron, params, phone, remark);
			if err != nil {
				result.Message = "更新失败,请重试!"
				this.WriteJson(result)

			} else {
				result.Success = true
				result.Message = "更新成功!";
				this.WriteJson(result)
			}
		}

	} else {
		// GET


		id, _ := this.GetInt("id")
		jobInfo, err := service.JobInfoService.FindJobInfoById(id)
		if err != nil {
			this.TplName = "500.html";


		} else {

			this.Data["jobInfo"] = jobInfo
			this.TplName = "jobinfo/edit.html";
		}

	}

}

// 详情
func (this *JobInfoController)Info() {

	id, _ := this.GetInt("id")
	jobInfo, err := service.JobInfoService.FindJobInfoById(id)
	if err != nil {
		this.TplName = "500.html";


	} else {

		this.Data["jobInfo"] = jobInfo
		this.TplName = "jobinfo/info.html";
	}

}
// 删除
func (this *JobInfoController)Delete() {
	result := common.Result{}

	id, _ := this.GetInt("Id");
	err := service.JobInfoService.DeleteJobInfoById(id)
	if err != nil {
		result.Message = "删除失败,请重试!"
	} else {
		result.Success = true
		result.Message = "删除成功"
	}

	this.WriteJson(result)
}

type HomeController struct {
	BaseController
}

func (this *HomeController) Index() {

	this.TplName = "index.html";
}

type JobSanpshotController struct {

  BaseController
}

// 查询任务执行快照列表
func (this *JobSanpshotController)List()  {

	sanpshotList,err := service.JobSanpshotService.List(0)
	if err!= nil {

		this.TplName = "500.html"
	} else {


		this.Data["sanpshotList"] = sanpshotList
		this.TplName = "jobsanpshot/list.html"
	}

}

func (this *JobSanpshotController)Info() {
	id,_ := this.GetInt("id")

	jobSnapshot,err := service.JobSanpshotService.FindJobSanpshotById(id,0)

	if err != nil {
		this.TplName = "500.html"
	}else {
		this.Data["jobSnapshot"] = jobSnapshot
		this.TplName = "jobsanpshot/info.html"
	}

}