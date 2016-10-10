package job

import (
	"github.com/shotdog/quartz"
	"kitty/app/service"
)

var jm *JobManager
type JobManager struct {
	qz *quartz.Quartz
}

func NewJobManager()*JobManager {

	if jm == nil {

		qz := quartz.New()
		qz.BootStrap()
		jm = &JobManager{qz:qz}
	}

	return jm
}

func (this *JobManager)PushAllJob()  {

	list,err := service.JobInfoService.List()
	if err!= nil || len(list) == 0 {
		return
	}


	



}