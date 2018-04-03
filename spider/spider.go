package spider

import (
	"github.com/liunian1004/go-crawler/page"
	"sync"
	"github.com/liunian1004/go-crawler/scheduler"
	"github.com/liunian1004/go-crawler/request"
)

type Spider struct {
	ID int
	Name string
	SubName string
	Description string
	UseCookie bool
	NotDefaultField bool // 是否禁止输出结果中的默认字段 Url/ParentUrl/DownloadTime

	RuleTree RuleTree

	Matrix *scheduler.Matrix // 存储请求

	Status scheduler.Status // 执行状态
	lock sync.RWMutex
	once sync.Once
}

var GlobalSpiders map[string]*Spider

func (s *Spider) Register() {
	GlobalSpiders[s.Name] = s
}

// 查找对应的规则
func (s *Spider) GetRule(ruleName string) (rule *Rule, r bool) {
	rule, r = s.RuleTree.Trunk[ruleName]
	return
}

func (s *Spider) Start() {
	s.RuleTree.Root(&Context{Spider: s})
	s.lock.Lock()
	s.Status = scheduler.RUN
	s.lock.Unlock()
}

func (s *Spider) Stop() {
	s.lock.Lock()
	s.Status = scheduler.STOP
	s.lock.Unlock()
}

// 初始化队列
func (s *Spider) SpiderMatrixInit() *Spider {
	s.Matrix = scheduler.InsertGlobalSchedulerMatrix(s.Name)
	return s
}

// 添加请求
func (s *Spider) PushRequest(request *request.Request) {
	s.Matrix.Push(request)
}

// 取出请求
func (s *Spider) PullRequest() (request *request.Request) {
	return s.Matrix.Pull()
}

// 保持成功记录
func (s *Spider) SaveSuccess() {
	
}

// 保存失败记录
func (s *Spider) SaveFailure() {

}

func (s *Spider) CanStop() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Status == scheduler.STOP
}

func (s *Spider) Parse(ruleName string, p *page.Page) *page.Page {
	// 根据 rule 将数据解析到 Items 和 Files
	// get Rule
	rule, _ := s.GetRule(ruleName)

	if rule.ParseFunc == nil {
		panic("解析函数不存在")
	}

	ctx := GetContext(s)
	// 解析数据
	rule.ParseFunc(ctx, p)

	PutContext(ctx)

	return p
}