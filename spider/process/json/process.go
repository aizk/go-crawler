package json

import (
	"Vua-Crawler/spider/model"
	"io/ioutil"
	"net/http"
)

type JSONProcess struct {
	jsonProcess *model.Process
}

func NewJSONProcess(jsonProcess *model.Process) *JSONProcess {
	return &JSONProcess{jsonProcess: jsonProcess}
}

func (j *JSONProcess) Process(response *http.Response) (*model.Page, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return
}
