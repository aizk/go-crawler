package scheduler

import (
	"github.com/liunian1004/go-crawler/request"
)

type Matrix struct {
	SpiderName string
	//MaxPage int64 // 最大采集页数
	Requests []*request.Request // 请求队列
}

func NewMatrix(spiderName string) *Matrix {
	return &Matrix{
		SpiderName: spiderName,
		//MaxPage: int64(maxPage),
		Requests: make([]*request.Request, 0),
	}
}

func (m *Matrix) Push(req *request.Request) {
	// 停止爬虫的请求上限
	//if m.MaxPage >= 0 {
	//	return
	//}

	// 根据 Scheduler 状态限制队列中的请求数

	// 去除重复的请求

	m.Requests = append(m.Requests, req)

	//atomic.AddInt64(&m.MaxPage, 1)
}

func (m *Matrix) Pull() (req *request.Request) {
	req = m.Requests[0]

	// 如果使用代理，设置代理。

	m.Requests = m.Requests[1:]
	return
}