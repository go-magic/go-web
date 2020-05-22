package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MySql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Sql      string `yaml:"sql"`
}

type Sys struct {
	Port string `yaml:"port"`
}

type Config struct {
	MySql
	Sys
}

var config *Config

func init() {
	config = new(Config)
	b, err := ioutil.ReadFile("./conf/config.yml")
	if err != nil {
		panic(err)
		return
	}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		panic(err)
		return
	}
}

func GetConfig() *Config {
	return config
}
