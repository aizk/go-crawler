package spider

import "sync"

type Context struct {
	Spider *Spider
}

var contextPool = &sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

func GetContext(spider *Spider) (ctx *Context) {
	ctx = contextPool.Get().(*Context)
	ctx.Spider = spider
	return
}

func PutContext(ctx *Context) {
	ctx.Spider = nil
	contextPool.Put(ctx)
}

// 生成下一步的函数，可以指定执行的 Rule Name 中的 GeneratorFunc
//func (c *Context) Generator() interface{} {
//
//	return
//}