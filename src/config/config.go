package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YAWAL/tcbe/src/db"
)

type Config struct {
	Port string         `json:"port"`
	DB   db.MongoConfig `json:"database"`
}

func LoadConfig(path string) (conf *Config, err error) {
	confData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(confData, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
