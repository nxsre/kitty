package service

import (
	"github.com/astaxie/beego/orm"
	"kitty/app/model"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ormer orm.Ormer
	JobInfoService *jobInfoService
	JobHistoryService *jobHistoryService
	JobSanpshotService *jobSanpshotService
	JobSanpshotHistoryService *jobSanpshotHistoryService
)

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/kitty?charset=utf8")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.RegisterModel(&model.JobInfo{}, &model.JobInfoHistory{}, &model.JobSnapshot{}, &model.JobSnapshotHistory{})
	ormer = orm.NewOrm()
	initService()

}
func initService() {
	JobInfoService = &jobInfoService{}
	JobHistoryService = &jobHistoryService{}
	JobSanpshotService = &jobSanpshotService{}
	JobSanpshotHistoryService = &jobSanpshotHistoryService{}

}
