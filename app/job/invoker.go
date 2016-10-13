package job

import (
	"github.com/shotdog/quartz"
	"kitty/app/service"
	"time"
	"kitty/app/model"
	"log"
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

func (this *jobManager)PushAllJob() {

	list, err := service.JobInfoService.List()
	if err != nil || len(list) == 0 {
		return
	}

	for _, jobInfo := range list {

		this.AddJob(jobInfo)

	}

}

func (this *jobManager)AddJob(jobInfo model.JobInfo) error {

	return this.qz.AddJob(&quartz.Job{
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
func (this *jobManager)ModifyJob(jobInfo *model.JobInfo) error {

	return this.qz.ModifyJob(&quartz.Job{
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
func (this *jobManager)RemoveJob(jobInfo model.JobInfo) error {

	return this.qz.RemoveJob(jobInfo.Id)

}

func invoke(jobId int, targetUrl, params string, nextTime time.Time) {
	jobInfo, err := service.JobInfoService.FindJobInfoById(jobId)
	if err != nil || jobInfo.Active == 0 {
		JobManager.RemoveJob(jobInfo)
		return
	}

	//

	initExecute(jobInfo,targetUrl,nextTime)


}

func initExecute(jobInfo model.JobInfo, targetUrl string, nextTime time.Time) {

	sanpshot:= &model.JobSnapshot{
		JobName:jobInfo.JobName,
		JobGroup:jobInfo.JobGroup,
		Cron:jobInfo.Cron,
		Url:jobInfo.Url,
		Detail:"【"+time.Now().Format("2006-01-02 15:04:05")+"】准备执行--->目标服务器地址:" + targetUrl,
		CreateTime:time.Now(),
		State:0,

	}


	err := service.JobSnapshotService.Add(sanpshot)
	if err!= nil {
		return
	}

	log.Println("snapshot:",sanpshot)


}

func invokeJob(snapshot *model.JobSnapshot)  {

	err := service.JobSnapshotService.Update(sanpshot)
	if err!= nil {
		return
	}

	log.Println("snapshot:",sanpshot)

}





type JobInvoker struct {

}



