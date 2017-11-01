package json

import (
	"Vua-Crawler/spider/logger"
	"Vua-Crawler/spider/model"
	"strings"

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

	var resultType string
	contentRoot := []string{}

	// contentRoot 为需要读取的节点名字，可以有多个
	if ok {
		contentInfo := strings.Split(v, "|")
		resultType = contentInfo[0]
		contentRoot = strings.Split(contentInfo[1], ".")
	}

	if resultType == "array" {
		// 读取字段
		result := []map[string]interface{}{}

		// 读取节点中的内容
		for _, name := range contentRoot {
			simpJSON = simpJSON.Get(name)
		}

		// 返回节点下的内容
		rootNode, err := simpJSON.Array()
		if err != nil {
			logger.Error("Json transform fail.")
			return nil, err
		}

		if len(rootNode) >= 0 {
			for _, node := range rootNode {
				nodeMap, ok := node.(map[string]interface{})
				if !ok {
					continue
				}
				data := map[string]interface{}{}
				for key, value := range jsonRule {
					if key == "node" {
						continue
					}
					data[key] = nodeMap[value]
				}
				result = append(result, data)
			}
		}

		page.Urls = []*model.Request{}
		page.Result = result
		page.ResultCount = len(result)
	}

	return page, nil
}
