package spider

import (
	"Vua-Crawler/spider/model"
)

type Spider struct {
	Id   string
	Name string

	Depth    int
	EndCount int

	Requests []*model.Request

	Process map[string]process.Process
	Pipline pipline.Pipline
}
