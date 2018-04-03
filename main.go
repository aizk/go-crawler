package main

import (
	"runtime"
	"github.com/liunian1004/go-crawler/crawler"
	"github.com/liunian1004/go-crawler/spider"
	"fmt"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	c := crawler.Crawler{}
	c.Init(&spider.Spider{
		Name: "Test",
		RuleTree: spider.RuleTree{
			Root: func(context *spider.Context) {
				fmt.Println("root")
			},
			Trunk: map[string]*spider.Rule {

			},
		},
	}).Run()
	c.Spider.Stop()
}