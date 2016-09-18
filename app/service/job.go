package service

import "kitty/app/model"

type jobInfoService struct {

}

func (this *jobInfoService)List() ([]model.JobInfo, error) {
	var infos []model.JobInfo
	_, err := ormer.QueryTable("job_info").OrderBy("-create_time").All(&infos)
	return infos, err
}

type jobHistoryService struct {

}

type jobSanpshotService struct {

}

type jobSanpshotHistoryService struct {

}

