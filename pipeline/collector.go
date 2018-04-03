package pipeline

// 数据收集器，实现 Pipeline 接口
type Collector struct {
	ItemChan chan Item
	FileChan chan File

	ItemCache []Item // 分配输出结果的缓存

	outputType string // 输出方式

	ItemCount int
	ItemOutputTimes int
	FileCount int

	Total int // 收集的数据总数
}

func NewCollector() *Collector {
	c := &Collector{}
	c.outputType = "console"
	c.ItemChan = make(chan Item)
	c.FileChan = make(chan File)
	c.ItemCache = make([]Item, 5)
	c.ItemCount = 0
	c.ItemOutputTimes = 0
	c.Total = 0
	return c
}

func (c *Collector) Start() {
	go func() {
		// 从队列迭代获取数据，直至队列 close
		for data := range c.ItemChan {
			c.ItemCache = append(c.ItemCache, data)

			// 数据条数
			c.ItemCount++
			
			// 判断是否满足一组数据的输出条件，或者是超时，先不实现
			if len(c.ItemCache) < 1 {
				continue
			}

			c.ItemOutputTimes++
			// 输出 阻塞调用
			c.outputData()
		}
		// Stop 退出时 close(ItemChan) 时调用，输出剩下的数据
		c.ItemOutputTimes++
		c.outputData()
	}()
}

func (c *Collector) Stop() {
	panic("implement me")
}

func (c *Collector) CollectFiles(files []File) error {
	for _, file := range files {
		// 写入 chan
		c.FileChan <- file
	}
	return nil
}

func (c *Collector) CollectItems(items []Item) error {
	for _, item := range items {
		c.ItemChan <- item
	}
	return nil
}

func (c *Collector) outputData()  {
	// 输出数据

}