package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var YamlConfigGlobal *YamlConfig

func NewYamlConfig(path string) (*YamlConfig, error) {
	yaml := &YamlConfig{
		path: path,
	}
	err := yaml.parse()
	return yaml, err
}

type YamlConfig struct {
	path    string
	Serverj string   `json:"serverj"`
	Juejin  []string `json:"juejin"`
}

// 解析 config
func (y *YamlConfig) parse() error {
	byteList, err := ioutil.ReadFile(y.path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(byteList, y)
}
