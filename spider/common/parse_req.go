package common

import (
	"vuacrawler/spider/model"
	"vuacrawler/spider/logger"
	"strings"
	"strconv"
)

// 解析请求
func ParseReq(r []*model.Request, query map[string]interface{}) []*model.Request {
	results := []*model.Request{}
	for _,req := range r {
		// 获取 rule
		res, ok := getRuleRequests(req, query)
		if ok {
			results = append(results, res...)
		} else {
			results = append(results, req)
		}
	}
	return results
}

// 解析规则语法的 URL
// Query string 
func getRuleRequests(r *model.Request, query map[string]interface{}) ([]*model.Request, bool) {
	reqs := []*model.Request{}

	// 查找匹配规则
	rule, match := getRuleAndMatch(r.Url)

	if !match {
		return nil, false
	}

	// if ctx != nil {
		
	// }

	reqs, match = ParseOffset(r, rule)

	return reqs, match
}

// 获取规则函数
func getRuleAndMatch(url string) (rule string, match bool) {
	find := false
	for i := 0; i < len(url); i++ {
		if find {
			if url[i] == '}' {
				match = true
				break
			}
			rule += string(url[i]) // 获取从 { 开始的每个字符直到 }
		}

		if url[i] == '{' {
			find = true
		}
	}
	return rule, match
}

// 解析带规则的 URL 为多个具体的 URL
// http://xxxxxxxx.com/abc/{begin-end,offset}/ example:{1-400,10}
// 转换后 http://xxxxxxxx.com/abc/11/
func ParseOffset(r *model.Request, rule string)	([]*model.Request, bool) {
	reqs := []*model.Request{}

	ruleSplit := strings.Split(rule, ",")
	if len(ruleSplit) != 2 {
		return reqs, false
	}

	beginEndSplit := strings.Split(ruleSplit[0], "-")
	
	var begin, end, offset int
	var err error

	begin, err = strconv.Atoi(beginEndSplit[0])
	end, err = strconv.Atoi(beginEndSplit[1])
	offset, err = strconv.Atoi(ruleSplit[1])

	if err != nil {
		return reqs, false
	}

	if offset == 0 {
		return reqs, false
	}

	logger.Log("begin: ", begin, "end: ", end, "offset: ", offset)
	for i := begin; i < end; i += offset {
		url := strings.Replace(r.Url, "{" + rule + "}", strconv.Itoa(i), 1)
		req := &model.Request{
			Url: url,
			Method: r.Method,
			ContentType: r.ContentType,
			Data: r.Data,
			Header: r.Header,
			Cookies: r.Cookies,
			ProcessName: r.ProcessName,
		}
		reqs = append(reqs, req)
	}

	return reqs, true
}