package crawler

import (
	"github.com/liunian1004/go-crawler/spider"
	"github.com/liunian1004/go-crawler/download"
	"github.com/liunian1004/go-crawler/pipeline"
	page2 "github.com/liunian1004/go-crawler/page"
	"time"
	"github.com/liunian1004/go-crawler/request"
)

// 采集器
type CrawlerInterface interface {
	Init(spider *spider.Spider) CrawlerInterface // 初始化
	Run()
	Stop()
	GetID() int
	CanStop() bool
}

// 从抽象层面调度各模块
type Crawler struct {
	Spider *spider.Spider // 采集蜘蛛
	Download download.Downloader // 下载器
	Pipeline pipeline.Pipeliner // 数据输出
	ID int // 采集器 ID
}

func New(id int) CrawlerInterface {
	return &Crawler{
		ID: id,
		Download: download.EasyDownload{},
	}
}

func (c *Crawler) Init(spider *spider.Spider) CrawlerInterface {
	// 初始化蜘蛛 scheduler

	// 初始化 Pipeline

	return c
}

func (c *Crawler) Run() {
	// 开启收集
	c.Pipeline.Start()

	t := make(chan bool)

	go func() {
		c.run()
		close(t)
	}()

	c.Spider.Start()

	// 阻塞等待退出
	<-t

	c.Pipeline.Stop()
}

func (c *Crawler) run() {
	for {
		// 从队列中取出一条请求
		req := c.GetOneRequest()
		if req == nil {
			// 是否可以停止任务
			//if self.Spider.CanStop() {
			//	break
			//}
			time.Sleep(20 * time.Millisecond)
			continue
		}

		// 执行请求，限制当前的并发数，获取一个资源，获取不到时阻塞
		c.UseOneThread()
		go func() {
			defer func() {
				c.ReleaseOneThread()
			}()
			// 拿到进入资源，进入执行抓取
			c.Process(req)
		}()

		// 等待一会
		time.Sleep(20 * time.Millisecond)
	}
}

func (c *Crawler) Process(req *request.Request) {
	// 下载的 URL

	// 下载的蜘蛛
	spider1 := c.Spider

	// 使用下载器下载
	page := c.Download.Download(spider1, req)

	// 使用规则，解析下载回来的 page 对象 (每个 Request 都有对应的规则)
	page.Parse(req.GetRuleName())

	err := c.Pipeline.CollectItems(page.GetItems())
	if err != nil {
		return
	}
	err = c.Pipeline.CollectFiles(page.GetFiles())
	if err != nil {
		return
	}

	// 记录成功的请求，判断是否作为失败的请求添加至队列尾部

	// 统计成功页数

	// 释放 Page
	page2.ReleasePage(page)
}

func (*Crawler) Stop() {
	panic("implement me")
}

func (c *Crawler) GetID() int {
	return c.ID
}

func (c *Crawler) SetID(id int) {
	c.ID = id
}

func (c *Crawler) CanStop() bool {
	panic("implement me")
}

func (c *Crawler) GetOneRequest() *request.Request {
	return nil
}

func (c *Crawler) UseOneThread() {
}

func (c *Crawler) ReleaseOneThread() {
}

// 辅助方法

