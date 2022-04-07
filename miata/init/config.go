package init

import (
	"io/ioutil"
	"errors"
	ioCheck "miata/util/io"
	jsonUtil "miata/util/json"
)

type Config struct {
	Zap Zap `json:"zap"`
	Mysql Mysql `json:"mysql"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if ioCheck.Exists("config.json") {
		buf, err := ioutil.ReadFile("config.json")
		if err != nil {
			return cfg, err
		}
		jsonUtil.FromJson(string(buf), &cfg)
	} else {
		return nil, errors.New("no config.json found")
	}
	return cfg, nil
}
