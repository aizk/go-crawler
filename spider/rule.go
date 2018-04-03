package spider

import "github.com/liunian1004/go-crawler/page"

// 采集规则节点
type Rule struct {
	ParseFunc func(page *page.Page)
}

type RuleTree struct {
	// 根节点，此时 Page 为空
	Root func(page *page.Page)
	Trunk map[string]*Rule
}