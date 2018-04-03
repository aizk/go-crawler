package download

import (
	"github.com/liunian1004/go-crawler/spider"
	"github.com/liunian1004/go-crawler/page"
	"github.com/liunian1004/go-crawler/pipeline"
	"github.com/liunian1004/go-crawler/request"
)

type Downloader interface{
	Download(spider *spider.Spider, request *request.Request) *page.Page // 根据抽象规则和请求对象下载页面
}

type EasyDownload struct {
}

func (EasyDownload) Download(spider *spider.Spider, request *request.Request) *page.Page {
	return &page.Page{
		Items: []pipeline.Item{
			map[string]interface{}{ "1": 1},
			},
	}
}
