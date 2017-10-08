package job

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shotdog/quartz"
	"github.com/soopsio/kitty/app/common"
	"github.com/soopsio/kitty/app/model"
	"github.com/soopsio/kitty/app/service"
)

var JobManager *jobManager

type jobManager struct {
	qz *quartz.Quartz
}

func NewJobManager() {

	if JobManager == nil {

		qz := quartz.New()
		qz.BootStrap()
		JobManager = &jobManager{qz: qz}
	}

}

func (this *jobManager) PushAllJob() {

	list, err := service.JobInfoService.List()
	if err != nil || len(list) == 0 {
		return
	}

	for _, jobInfo := range list {

		this.AddJob(jobInfo)

	}

}

func (this *jobManager) AddJob(jobInfo model.JobInfo) error {

	return this.qz.AddJob(&quartz.Job{
		Id:         jobInfo.Id,
		Name:       jobInfo.JobName,
		Group:      jobInfo.JobGroup,
		Url:        jobInfo.Url,
		Params:     jobInfo.Params,
		Expression: jobInfo.Cron,
		JobFunc:    invoke,
	})

}

// modify
func (this *jobManager) ModifyJob(jobInfo *model.JobInfo) error {

	return this.qz.ModifyJob(&quartz.Job{
		Id:         jobInfo.Id,
		Name:       jobInfo.JobName,
		Group:      jobInfo.JobGroup,
		Url:        jobInfo.Url,
		Params:     jobInfo.Params,
		Expression: jobInfo.Cron,
		JobFunc:    invoke,
	})
}

// remove
func (this *jobManager) RemoveJob(jobInfo model.JobInfo) error {

	return this.qz.RemoveJob(jobInfo.Id)

}

func (this *jobManager) List() ([]*quartz.Job, error) {

	return this.qz.SnapshotJob()
}

func invoke(jobId int, targetUrl, params string, nextTime time.Time) {
	jobInfo, err := service.JobInfoService.FindJobInfoById(jobId)
	if err != nil || jobInfo.Active == 0 {
		JobManager.RemoveJob(jobInfo)
		return
	}

	initExecute(jobInfo, targetUrl, nextTime)

}

// 1.初始化
func initExecute(jobInfo model.JobInfo, targetUrl string, nextTime time.Time) {

	snapshot := &model.JobSnapshot{
		JobName:    jobInfo.JobName,
		JobGroup:   jobInfo.JobGroup,
		Cron:       jobInfo.Cron,
		Url:        jobInfo.Url,
		Params:     jobInfo.Params,
		JobId:      jobInfo.Id,
		Detail:     "【" + time.Now().Format("2006-01-02 15:04:05") + "】初始化完成目标服务器:" + targetUrl,
		CreateTime: time.Now(),
		State:      0,
	}

	err := service.JobSnapshotService.Add(snapshot)
	if err != nil {
		return
	}

	log.Println("snapshot:", snapshot)

	invokeJob(snapshot)

}

func invokeJob(snapshot *model.JobSnapshot) {

	detail := snapshot.Detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】正在调用..."
	err := service.JobSnapshotService.Update(snapshot.Id, 1, detail, time.Now(), "", 0)
	if err != nil {
		return
	}

	req := common.Request{
		SnapshotId: snapshot.Id,
		JobId:      snapshot.JobId,
		Params:     snapshot.Params,
		Method:     "INVOKE",
	}
	body, _ := json.Marshal(req)
	res, err := http.Post(snapshot.Url, "application/json;charset=utf-8", bytes.NewReader(body))
	if err != nil {

		detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器不可用..." + err.Error()
		service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
		return

	} else {

		bys := make([]byte, 1024)
		n, _ := res.Body.Read(bys)

		body := common.Response{}
		err = json.Unmarshal(bys[:n], &body)
		if err != nil {
			detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器不可用..." + err.Error()
			service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
			return
		} else {

			if body.State == 4 {
				detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器任务执行失败..." + err.Error()
				service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
				return
			} else if body.State == 2 {

				detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器任务正在执行..." + err.Error()
				service.JobSnapshotService.Update(snapshot.Id, 2, detail, time.Now(), "", 0)
				snapshot.Detail = detail
				checkExecutionJob(snapshot)

			}

		}

	}

}

//检查执行job情况
func checkExecutionJob(snapshot *model.JobSnapshot) {

	for {
		select {
		case <-time.After(time.Second * 10):
			log.Println("**********任务检查**********", snapshot.Id, snapshot.JobId)
			req := common.Request{
				SnapshotId: snapshot.Id,
				JobId:      snapshot.JobId,
				Params:     snapshot.Params,
				Method:     "CHECK",
			}

			body, _ := json.Marshal(req)
			detail := snapshot.Detail
			res, err := http.Post(snapshot.Url, "application/json;charset=utf-8", bytes.NewReader(body))
			if err != nil {

				detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器不可用..." + err.Error()
				service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
				return

			} else {

				bys := make([]byte, 1024)
				n, _ := res.Body.Read(bys)

				response := common.Response{}
				err = json.Unmarshal(bys[:n], &response)
				if err != nil {
					detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器不可用..." + err.Error()
					service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
					return
				} else {

					if response.State == 3 {
						detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器任务已完成..."
						service.JobSnapshotService.Update(snapshot.Id, 3, detail, time.Now(), response.Result, response.TimeConsume)
						return
					} else if response.State == 4 {

						detail = detail + "\n【" + time.Now().Format("2006-01-02 15:04:05") + "】目标服务器任务执行失败..."
						service.JobSnapshotService.Update(snapshot.Id, 4, detail, time.Now(), "", 0)
						snapshot.Detail = detail

					}

				}

			}
		}
	}

}

type JobInvoker struct {
}
