package json

import (
	"Vua-Crawler/spider/logger"
	"Vua-Crawler/spider/model"

	simplejson "github.com/bitly/go-simplejson"
)

// 解析规则
func Rule(process *model.Process, bytes []byte) (*model.Page, error) {
	page := &model.Page{}

	simpJSON, err := simplejson.NewJson(bytes)
	if err != nil {
		logger.Error("NewJson fail in json rule.go", err)
		return nil, err
	}

	jsonRule := process.JsonRule.Rule

	v, ok := jsonRule["node"]

	if ok {

	}
}
