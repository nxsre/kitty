package job

import (
	"github.com/shotdog/quartz"
	"kitty/app/service"
	"time"
	"kitty/app/model"
)

var JobManager *jobManager
type jobManager struct {
	qz *quartz.Quartz
}

func NewJobManager() {

	if JobManager == nil {

		qz := quartz.New()
		qz.BootStrap()
		JobManager = &jobManager{qz:qz}
	}


}

func (this *jobManager)PushAllJob()  {

	list,err := service.JobInfoService.List()
	if err!= nil || len(list) == 0 {
		return
	}

	for _,jobInfo:=range list{


		this.AddJob(jobInfo)

	}
	

}

func (this *jobManager)AddJob(jobInfo model.JobInfo) error {

	return  this.qz.AddJob(&quartz.Job{
		Id:jobInfo.Id,
		Name:jobInfo.JobName,
		Group:jobInfo.JobGroup,
		Url:jobInfo.Url,
		Params:jobInfo.Params,
		Expression:jobInfo.Cron,
		JobFunc:invoke,

	})

}

// modify
func (this *jobManager)ModifyJob(jobInfo *model.JobInfo)error  {

	return  this.qz.ModifyJob(&quartz.Job{
		Id:jobInfo.Id,
		Name:jobInfo.JobName,
		Group:jobInfo.JobGroup,
		Url:jobInfo.Url,
		Params:jobInfo.Params,
		Expression:jobInfo.Cron,
		JobFunc:invoke,

	})
}

// remove
func (this *jobManager)RemoveJob(jobInfo model.JobInfo)error  {

	return  this.qz.RemoveJob(jobInfo.Id)

}



func invoke(jobId int, targetUrl, params string, nextTime time.Time)  {
	jobInfo ,err := service.JobInfoService.FindJobInfoById(jobId)
	if err != nil || jobInfo.Active == 0{
		JobManager.RemoveJob(jobInfo)
		return
	}

	//




}

type JobInvoker struct {

}