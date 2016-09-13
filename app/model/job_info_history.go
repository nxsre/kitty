package model

import "time"


type JobInfoHistory struct {
	Id         int       // 主键
	JobName    string    // 任务名称
	JobGroup   string    // 任务分组
	Params     string    // 任务参数
	Cron       string    // cron 表达式
	Url        string    // 任务目标地址
	Phone      string    // 任务负责人手机号码
	Active     int       // 是否为激活状态
	Remark     string    // 备注
	CreateTime time.Time // 任务创建时间
	UpdateTime time.Time // 任务更新时间
}