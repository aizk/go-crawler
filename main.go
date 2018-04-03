package main

import (
	"runtime"
	"github.com/liunian1004/go-crawler/crawler"
	"github.com/liunian1004/go-crawler/spider"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	c := crawler.Crawler{}
	c.Init(&spider.Spider{
		Name: "Test",
	}).Run()
}