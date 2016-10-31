package controller

import (
	"kitty/app/service"
	"github.com/astaxie/beego/validation"
	"kitty/app/common"
	"kitty/app/job"
	"kitty/app/model"
)

type JobInfoController struct {
	BaseController
}

// 查询任务列表
func (this *JobInfoController)List() {

	pageNo, _ := this.GetInt("pageNo", 1)
	pageSize, _ := this.GetInt("pageSize", 10)

	jobName := this.GetString("JobName")
	groupName := this.GetString("GroupName")
	infos, _ := service.JobInfoService.FindJobInfoListByPage(pageNo, pageSize, jobName, groupName)

	count, _ := service.JobInfoService.FindJobInfoCountByState(0, jobName, groupName);


	pager := &common.Pager{PageNo:pageNo, PageSize:pageSize}
	pager.SetTotalCount(count)
	pager.Data = infos

	this.Data["pager"] = pager
	this.Data["jobName"] = jobName
	this.Data["groupName"] = groupName
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

			info, err := service.JobInfoService.FindJobInfoById(id)
			if err != nil {
				result.Message = "此任务不存在!";
				this.WriteJson(result)
			}

			info.Cron = cron
			info.Params = params
			info.Phone = phone
			info.Url = url
			info.Remark = remark
			job.JobManager.ModifyJob(&info)

			err = service.JobInfoService.UpdateJobInfo(id, url, cron, params, phone, remark);
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

	jobInfo, err := service.JobInfoService.FindJobInfoById(id)

	if err != nil {
		result.Message = "此任务不存在!"
	} else {

		job.JobManager.RemoveJob(jobInfo)
		err = service.JobInfoService.DeleteJobInfoById(id)
		if err != nil {
			result.Message = "删除失败,请重试!"
		} else {
			result.Success = true
			result.Message = "删除成功"
		}

	}
	this.WriteJson(result)
}


// 激活或者取消激活

func (this *JobInfoController)Active() {
	result := common.Result{}
	active, _ := this.GetInt("active")
	id, _ := this.GetInt("id")

	jobInfo, err := service.JobInfoService.FindJobInfoById(id)
	if err != nil {

		job.JobManager.RemoveJob(model.JobInfo{Id:id})
		result.Message = "此任务不存在!";
	} else {

		if active == 1 {

			jobInfo.Active = 1
			err = service.JobInfoService.UpdateJobActive(id, active)
			if err == nil {
				job.JobManager.AddJob(jobInfo)
				result.Message = "成功"
				result.Success = true
			} else {
				result.Message = "激活失败请重试!";

			}

		} else {
			jobInfo.Active = 0

			err = service.JobInfoService.UpdateJobActive(id, active)
			if err == nil {
				job.JobManager.RemoveJob(model.JobInfo{Id:id})

				result.Message = "成功"
				result.Success = true
			} else {
				result.Message = "取消激活失败请重试!";

			}

		}

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
func (this *JobSanpshotController)List() {

	pageNo, _ := this.GetInt("pageNo", 1)
	pageSize, _ := this.GetInt("pageSize", 10)

	jobName := this.GetString("JobName")
	groupName := this.GetString("GroupName")
	state, _ := this.GetInt("State", -1)

	snapshotList, err := service.JobSnapshotService.FindJobSnapshotInfoListByPage(pageNo, pageSize, jobName, groupName, state)

	count, _ := service.JobSnapshotService.FindJobSnapshotCount(jobName, groupName, state);
	if err != nil {

		this.TplName = "500.html"
	} else {

		pager := &common.Pager{PageNo:pageNo, PageSize:pageSize}
		pager.SetTotalCount(count)
		pager.Data = snapshotList
		this.Data["pager"] = pager
		this.Data["jobName"] = jobName
		this.Data["groupName"] = groupName
		this.Data["sanpshotList"] = snapshotList
		this.Data["State"] = state
		this.TplName = "jobsnapshot/list.html"
	}

}



// 查询详情
func (this *JobSanpshotController)Info() {
	id, _ := this.GetInt("id")

	jobSnapshot, err := service.JobSnapshotService.FindJobSnapshotById(id, 0)

	if err != nil {
		this.TplName = "500.html"
	} else {
		this.Data["jobSnapshot"] = jobSnapshot
		this.TplName = "jobsanpshot/info.html"
	}

}

func (this *JobSanpshotController)Delete() {

	result := common.Result{}
	id, _ := this.GetInt("Id", -1)

	if id == -1 {
		result.Message = "此任务不存在!"
		this.WriteJson(result)
	}

	_, err := service.JobSnapshotService.FindJobSnapshotById(id, 0)
	if err != nil  {
		result.Message = "此任务不存在!"
		this.WriteJson(result)

	}

	err = service.JobSnapshotService.DeleteJobSnapshotById(id);

	if err != nil {

		result.Message = "此任务不存在!"
		this.WriteJson(result)
	} else {

		result.Message = "删除成功"
		result.Success = true
		this.WriteJson(result)

	}

}

type MonitorController struct {
	BaseController
}

// 监控
func (this *MonitorController)List() {

	list, _ := job.JobManager.List()

	this.Data["list"] = list
	this.TplName = "monitor/monitor.html"

}