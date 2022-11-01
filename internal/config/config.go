package config

import (
	"encoding/json"
	"io/ioutil"
)

const configPath = "config.json"

type Patterns struct {
	CreatePattern string `json:"createPattern"`
	UpdatePattern string `json:"updatePattern"`
	DeletePattern string `json:"deletePattern"`
	DayPattern    string `json:"dayPattern"`
	WeekPattern   string `json:"weekPattern"`
	MonthPattern  string `json:"monthPattern"`
}

type Config struct {
	Port         string   `json:"port"`
	DatabaseConn string   `json:"databaseConn"`
	Patterns     Patterns `json:"patterns"`
}

// InitConfig чтение конфига
func InitConfig() Config {
	var config Config

	if data, err := ioutil.ReadFile(configPath); err == nil {
		if err := json.Unmarshal(data, &config); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return config
}
