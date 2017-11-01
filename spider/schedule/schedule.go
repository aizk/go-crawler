package schedule

import (
	"Vua-Crawler/spider/common"
	"Vua-Crawler/spider/logger"
	"Vua-Crawler/spider/model"
)

// 任务队列
type Schedule struct {
	waitQueue chan *model.Request
}

// return NewSchedule
func NewSchedule(maxWaitNum int) *Schedule {
	schedule := &Schedule{}
	schedule.waitQueue = make(chan *model.Request, maxWaitNum) // 缓存chan
	return schedule
}

// push in queue
func (s *Schedule) Push(req *model.Request) {
	praseReqs := common.ParseReq([]*model.Request{req}, nil)
	for _, req := range praseReqs {
		logger.Log("Push Url: ", req.Url, req.ProcessName, len(s.waitQueue))
		s.waitQueue <- req
	}
}

func (s *Schedule) PushMuti(reqs []*model.Request) {
	praseReqs := common.ParseReq(reqs, nil)
	for _, req := range praseReqs {
		logger.Log("Push Url: ", req.Url, req.ProcessName, len(s.waitQueue))
		s.waitQueue <- req
	}
}

func (s *Schedule) Pop() (*model.Request, bool) {
	req, ok := <-s.waitQueue
	return req, ok
}

func (s *Schedule) Count() int {
	return len(s.waitQueue)
}

func (s *Schedule) Close() {
	close(s.waitQueue)
}
