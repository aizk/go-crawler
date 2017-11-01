package config

import (
	"Vua-Crawler/spider/logger"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	WorkNum    int      `json:"work_num"`
	MaxWaitNum int      `json:"max_wait_num"`
	HttpAddr   string   `json:"http_addr"`
	Etcd       []string `json:"etcd"`
}

var ConfigGlobal *Config

func InitConfig() error {
	var file *os.File
	var bytes []byte
	var err error

	file, err = os.OpenFile("conf.json", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	ConfigGlobal = &Config{}
	err = json.Unmarshal(bytes, ConfigGlobal)
	if err != nil {
		return err
	}

	logger.Log("Init success!", *ConfigGlobal)
	return nil
}
