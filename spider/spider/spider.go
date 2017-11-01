package spider

import (
	"Vua-Crawler/spider/model"
	"Vua-Crawler/spider/pipline"
	"Vua-Crawler/spider/pipline/console"
	"Vua-Crawler/spider/process"
	"Vua-Crawler/spider/process/json"
)

type Spider struct {
	ID   string
	Name string

	Depth    int
	EndCount int

	Requests []*model.Request

	Process map[string]process.Process
	Pipline pipline.Pipline
}

func InitForTask(task *model.Task) *Spider {
	s := &Spider{}
	s.ID = task.ID
	s.Name = task.Name
	s.Depth = task.Depth
	s.EndCount = task.EndCount
	s.Requests = task.Request

	s.Process = make(map[string]process.Process)

	// 根据配置初始化处理器
	for i, p := range task.Process {
		switch p.Type {
		case "json":
			s.Process[p.Name] = json.NewJSONProcess(&task.Process[i])
		}
	}

	switch task.Pipline {
	case "console":
		s.Pipline = console.NewConsolePipline()
	}

	return s
}

func (s *Spider) GetPipline() pipline.Pipline {
	return s.Pipline
}
