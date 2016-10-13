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
	JobSnapshotService *jobSnapshotService
	JobSnapshotHistoryService *jobSnapshotHistoryService
)

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/kitty?charset=utf8")
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.Debug = true
	orm.RegisterModel(new(model.JobInfo), new(model.JobInfoHistory), new(model.JobSnapshot), new(model.JobSnapshotHistory))
	ormer = orm.NewOrm()
	initService()

}
func initService() {
	JobInfoService = &jobInfoService{}
	JobHistoryService = &jobHistoryService{}
	JobSnapshotService = &jobSnapshotService{}
	JobSnapshotHistoryService = &jobSnapshotHistoryService{}

}
