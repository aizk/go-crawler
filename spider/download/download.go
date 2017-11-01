package download

import (
	"bytes"
	"net/http"
)

// 开发多种 Client
// 获取页面
func GetPage(id string, method string, url string, data []byte) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
