package crawler

import (
	"testing"
	"github.com/liunian1004/go-crawler/request"
	"github.com/liunian1004/go-crawler/spider"
	"fmt"
	"github.com/liunian1004/go-crawler/page"
)

func TestCrawler_Process(t *testing.T) {
	c := New(1)
	c.Init(&spider.Spider{
		Name: "Test",
		RuleTree: spider.RuleTree{
			Root: func(context *spider.Context) {
				fmt.Println("root")
			},
			Trunk: map[string]*spider.Rule {
				"123": {
					ParseFunc: func(ctx *spider.Context, page *page.Page) {
						fmt.Println("12345")
					},
				},
			},
		},
	})
	c.(*Crawler).Process(&request.Request{
		Rule: "123",
	})
}