package service

import (
	"kitty/app/model"
)

type jobInfoService struct {

}

func (this *jobInfoService)List() ([]model.JobInfo, error) {
	var infos []model.JobInfo
	_, err := ormer.QueryTable("job_info").OrderBy("-create_time").All(&infos)
	return infos, err
}

// 新增
func (this *jobInfoService)Add(jobName, jobGroup, cron, url, params, phone, remark string, active int) error {

	jobInfo := &model.JobInfo{
		JobName:jobName,
		JobGroup:jobGroup,
		Cron:cron,
		Url:url,
		Active:active,
		Params:params,
		Phone:phone,
		Remark:remark,
	}

	_, err := ormer.Insert(jobInfo)
	return err

}

func (this *jobInfoService)FindJobInfoById(id int) (model.JobInfo, error) {

	var jobInfo model.JobInfo
	err := ormer.QueryTable("job_info").Filter("id", id).One(&jobInfo)
	return jobInfo, err

}

func (this *jobInfoService)UpdateJobInfo(id int, url, cron, params, phone, remark string) error {
	jobInfo := &model.JobInfo{Id:id, Url:url, Cron:cron, Params:params, Phone:phone, Remark:remark}
	_, err := ormer.Update(jobInfo, "url", "cron", "params", "phone", "remark")
	return err

}

type jobHistoryService struct {

}

type jobSanpshotService struct {

}

type jobSanpshotHistoryService struct {

}

