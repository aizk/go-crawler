package scheduler

type Status int

// 运行状态
const (
	STOPPED Status = iota - 1
	STOP
	RUN
	PAUSE
)