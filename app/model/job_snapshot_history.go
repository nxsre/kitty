package model

// 任务快照历史
type JobSnapshotHistory struct {
	Id          int    // 主键
	JobName     string // 任务名称
	JobGroup    string // 任务分组
	Params      string // 参数
	Cron        string // cron 表达式
	Url         string // 目标服务器url地址
	Detail      string // 执行详情
	Ip          string // ip地址
	State       int    // 执行状态
	Result      string // 执行结果
	TimeConsume int64  // 任务耗时
}


