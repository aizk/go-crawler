package scheduler

import (
	"sync"
)

type Scheduler struct {
	Status      Status    // 运行状态
	ThreadCount chan bool // 并发数
	UseProxy    bool      // 是否使用代理
	//Proxy *Proxy
	Matrices []*Matrix
	sync.RWMutex
}

var GlobalScheduler = &Scheduler{
	Status: RUN,
	ThreadCount: make(chan bool, 10),
	//Proxy
	Matrices: []*Matrix{},
}

func InsertGlobalSchedulerMatrix(spiderName string) *Matrix {
	matrix := NewMatrix(spiderName)
	GlobalScheduler.RLock()
 	defer GlobalScheduler.RUnlock()
	GlobalScheduler.Matrices = append(GlobalScheduler.Matrices, matrix)
	return matrix
}