package console

import (
	"Vua-Crawler/spider/logger"
	"encoding/json"
)

type ConsolePipline struct {
}

func NewConsolePipline() *ConsolePipline {
	return &ConsolePipline{}
}

func (c *ConsolePipline) ProcessData(v interface{}, task string, process string) {
	bytes, _ := json.Marshal(v)
	logger.Log("Pipline: ", string(bytes))
}
