package common

type Request struct {
	SnapshotId int
	JobId      int
	Params     string
	Method     string
}

type Response struct {
	SnapshotId  int
	JobId       int
	Params      string
	Method      string
	State       int
	TimeConsume int64
	Result      string
}


