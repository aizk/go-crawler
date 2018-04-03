package pipeline

type Pipeliner interface{
	Start() // 启动收集
	Stop() // 停止收集
	CollectFiles(files []File) error
	CollectItems(items []Item) error
}

func NewPipeline() Pipeliner {
	return NewCollector()
}