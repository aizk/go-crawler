package spider

import "github.com/liunian1004/go-crawler/page"

// 采集规则节点
type Rule struct {
	ParseFunc func(ctx *Context, page *page.Page)
	GeneratorFunc func(ctx *Context, param map[string]interface{})
}

type RuleTree struct {
	// 根节点，此时 Page 为空
	Root func(ctx *Context)
	Trunk map[string]*Rule
}