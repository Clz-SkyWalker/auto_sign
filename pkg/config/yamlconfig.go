package config

import (
	"auto_sign/pkg/utils"
	"io/ioutil"

	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

var YamlConfigGlobal *YamlConfig

func NewYamlConfig(path string) (*YamlConfig, error) {
	yaml := &YamlConfig{
		path: path,
	}
	err := yaml.parse()
	if err != nil {
		utils.AddLogger(utils.ErrYamlUnmarshal, zapcore.Field{Key: "path", String: path})
	}
	return yaml, err
}

type YamlConfig struct {
	path    string
	Serverj string   `json:"serverj"`
	Juejin  []string `json:"juejin"`
	Ireader []string `json:"ireader"`
}

// 解析 config
func (y *YamlConfig) parse() error {
	byteList, err := ioutil.ReadFile(y.path)
	if err != nil {
		utils.AddLogger(utils.ErrReadFile, zapcore.Field{Key: "path", String: y.path})
		return err
	}
	return yaml.Unmarshal(byteList, y)
}
