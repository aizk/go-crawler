package core

import (
	"Vua-Crawler/spider/config"
	"Vua-Crawler/spider/download"
	"Vua-Crawler/spider/logger"
	"Vua-Crawler/spider/model"
	"Vua-Crawler/spider/schedule"
	"Vua-Crawler/spider/spider"
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const DefaultWorkNum = 1

type SpiderRuntime struct {
	sync.Mutex
	workNum  int
	schedule *schedule.Schedule
	spider   *spider.Spider

	stopSign    bool
	recoverChan chan int

	TaskMeta *TaskMeta
}

// TaskMeta
type TaskMeta struct {
	DownloadFailCount int32 `json:"download_fail_count"`
	DownloadCount     int32 `json:"download_count"`

	URLNum           int32 `json:"url_num"`
	WaitURLNum       int   `json:"wait_url_num"`
	CrawlerResultNum int32 `json:"crawler_result_num"`
}

func NewSpiderRuntime() *SpiderRuntime {
	workNum := config.ConfigGlobal.WorkNum
	if workNum == 0 {
		workNum = DefaultWorkNum
	}

	s := &SpiderRuntime{}
	s.workNum = workNum
	s.schedule = schedule.NewSchedule(config.ConfigGlobal.MaxWaitNum)
	s.recoverChan = make(chan int)

	meta := &TaskMeta{}
	meta.WaitURLNum = 0
	meta.URLNum = (int32)(0)
	meta.DownloadCount = (int32)(0)
	meta.DownloadFailCount = (int32)(0)
	meta.CrawlerResultNum = (int32)(0)

	s.TaskMeta = meta

	return s
}

func (s *SpiderRuntime) SetSpider(spider *spider.Spider) {
	s.spider = spider
}

func (s *SpiderRuntime) GetSpider() *spider.Spider {
	return s.spider
}

func (s *SpiderRuntime) Run() {
	if s.stopSign {
		s.recoverChan <- 1
		return
	}
	s.schedule.PushMuti(s.spider.Requests)

	// 执行
	for i := 0; i < s.workNum; i++ {
		go s.worker()
	}
}

func (s *SpiderRuntime) Stop() {
	s.stopSign = true
}

func (s *SpiderRuntime) Exit() {
	s.schedule.Close()
	close(s.recoverChan)
}

// 工作 goroutine
func (s *SpiderRuntime) worker() {
	for {
		// 如果 recoverChan 读取成功，则不退出
		if s.stopSign {
			_, ok := <-s.recoverChan
			s.stopSign = false
			if !ok {
				goto exit
			}
		}

		req, ok := s.schedule.Pop()
		if !ok {
			goto exit
		}

		atomic.AddInt32(&s.TaskMeta.DownloadCount, 1)
		response, err := s.download(req)
		if err != nil {
			logger.Error("spiderRuntime download fail.")
			atomic.AddInt32(&s.TaskMeta.DownloadFailCount, 1)
			continue
		}
		// 调用相应的 Process
		process, ok := s.spider.Process[req.ProcessName]
		if !ok {
			logger.Error("spiderRuntime get Process fail.")
			break
		}

		page, err := process.Process(response)
		if err != nil {
			logger.Error("spiderRuntime Process response fail.")
			break
		}

		atomic.AddInt32(&s.TaskMeta.URLNum, int32(len(page.Urls)))

		s.TaskMeta.WaitURLNum = s.schedule.Count()

		go func() {
			s.schedule.PushMuti(page.Urls)
		}()

		atomic.AddInt32(&s.TaskMeta.CrawlerResultNum, int32(page.ResultCount))

		s.spider.Pipline.ProcessData(page.Result, s.spider.Name, req.ProcessName)
	}
exit:
	logger.Log(s.spider.Name, "Worker Close.")
}

func (s *SpiderRuntime) download(req *model.Request) (*http.Response, error) {
	time.Sleep(1 * time.Second)
	// switch req.Method {
	// case "get":
	// 	return download.GetPage()
	// }
	data, err := json.Marshal(req.Data)
	if err != nil {
		logger.Error("spiderRuntim parse req.Data to []byte fail.")
		return nil, err
	}
	resp, err := download.GetPage("0", req.Method, req.Url, data)
	if err != nil {
		logger.Error("spiderRuntime download fail.")
		return nil, err
	}
	return resp, nil
}
