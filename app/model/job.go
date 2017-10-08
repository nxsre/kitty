package model

import "time"

type JobInfo struct {
	Id         int       // 主键
	JobName    string    // 任务名称
	JobGroup   string    // 任务分组
	State      int       // 状态  0 正常 1 删除
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
	State      int       // 状态  0 正常 1 删除
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
type JobSnapshot struct {
	Id          int       // 主键
	JobId       int       // 任务id
	JobName     string    // 任务名称
	JobGroup    string    //  任务分组
	Params      string    // 参数
	Cron        string    // cron表达式
	Url         string    // 目标服务器url
	Detail      string    // 执行详情
	Ip          string    // ip地址
	State       int       // 执行状态
	Result      string    // 执行结果
	TimeConsume int64     // 任务耗时
	CreateTime  time.Time // 任务创建时间
	UpdateTime  time.Time // 任务更新时间
}

// 任务快照历史
type JobSnapshotHistory struct {
	Id          int       // 主键
	JobId       int       // 任务id
	JobName     string    // 任务名称
	JobGroup    string    // 任务分组
	Params      string    // 参数
	Cron        string    // cron 表达式
	Url         string    // 目标服务器url地址
	Detail      string    // 执行详情
	Ip          string    // ip地址
	State       int       // 执行状态
	Result      string    // 执行结果
	TimeConsume int64     // 任务耗时
	CreateTime  time.Time // 任务创建时间
	UpdateTime  time.Time // 任务更新时间
}
