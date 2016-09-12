package model

import "time"

type JobInfo struct {
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

// 任务快照
type JobSapshot struct {
	Id          int
	JobName     string
	JobGroup    string
	Params      string
	Cron        string
	Url         string
	Detail      string
	Ip          string
	State       int
	Result      string
	TimeConsume int64
}


type JobSapshotHistory struct {
	Id          int
	JobName     string
	JobGroup    string
	Params      string
	Cron        string
	Url         string
	Detail      string
	Ip          string
	State       int
	Result      string
	TimeConsume int64
}


