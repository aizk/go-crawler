package request

type Request struct {
	Rule string // 解析响应的对应规则
}

func (r *Request) GetRuleName() string {
	return r.Rule
}