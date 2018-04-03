package page

import (
	"github.com/liunian1004/go-crawler/pipeline"
	"net/http"
	"github.com/liunian1004/go-crawler/request"
	//"github.com/liunian1004/go-crawler/spider"
	"sync"
)

type Pager interface{
	Parse(ruleName string) // 解析返回对象中的文件和数据
	GetItems()
	GetFiles()
}

type Page struct {
	//Spider *spider.Spider // 规则

	Request *request.Request // 请求
	Response *http.Response // 响应流

	Items []pipeline.Item
	Files []pipeline.File
}

var pagePool = &sync.Pool{
	New: func() interface{} {
		return &Page{
			Items: []pipeline.Item{},
			Files: []pipeline.File{},
		}
	},
}

func GetPage(request *request.Request) *Page {
	page := pagePool.Get().(*Page)
	//page.Spider = s
	page.Request = request
	return page
}

func ReleasePage(page *Page) {
	if page.Response != nil {
		page.Response.Body.Close()
		page.Response = nil
	}
	page.Request = nil
	//page.Spider = nil
	page.Items = page.Items[:0]
	page.Files = page.Files[:0]
	// 后续还要增加释放的资源
	pagePool.Put(page)
}

func (p *Page) GetItems() []pipeline.Item {
	return p.Items
}

func (p *Page) GetFiles() []pipeline.File {
	return p.Files
}

//func (p *Page) Parse(ruleName string) *Page {
//	// 根据 rule 将数据解析到 Items 和 Files
//	// get Rule
//	rule := p.getRule(ruleName)
//
//	if rule.ParseFunc == nil {
//		panic("解析函数不存在")
//	}
//
//	// 解析数据
//	rule.ParseFunc(p)
//
//	return p
//}

//func (p *Page) getRule(ruleName string) (rule *spider.Rule) {
//	rule, b := p.Spider.GetRule(ruleName)
//	if b == false {
//		panic("获取规则失败")
//	}
//	return
//}