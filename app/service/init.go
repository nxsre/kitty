package service

import (
	"github.com/astaxie/beego/orm"
	"kitty/app/model"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
	JobInfoService *jobInfoService
)

func Init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/kitty?charset=utf8")
	orm.RegisterModel(&model.JobInfo{}, &model.JobSnapshot{}, &model.JobInfoHistory{}, &model.JobSnapshotHistory{})
	// 设置最大空闲的连接
	orm.SetMaxIdleConns("default", 20)
	orm.SetMaxOpenConns("default", 100)
	orm.Debug = true
	o = orm.NewOrm()
}

func initService() {

	JobInfoService = &jobInfoService{}
}