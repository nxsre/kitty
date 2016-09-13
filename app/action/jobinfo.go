package action

import (
	"github.com/astaxie/beego"
	"kitty/app/service"
	"log"
)

type JobInfo struct {

	beego.Controller
}

func (this *JobInfo)Index()  {

	this.TplName = "job/job_info.html"
	jobInfoList,_ := service.JobInfoService.List()

	log.Println(jobInfoList)
	this.Data["jobInfoList"] = jobInfoList

}