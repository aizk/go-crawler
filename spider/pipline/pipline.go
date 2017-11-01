package pipline

type Pipline interface {
	ProcessData(v interface{}, task string, process string)
}
