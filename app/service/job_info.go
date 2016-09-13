package service

import (

	"kitty/app/model"
)

type jobInfoService struct {

}

func (this *jobInfoService)List() ([]model.JobInfo, error) {

	var jobInfoList []model.JobInfo
	_, err := o.QueryTable("job_info").All(&jobInfoList)
	return jobInfoList, err

}






