package service

import (
	"auto_sign/pkg/config"
	"auto_sign/pkg/utils"
)

func testInit() (*config.YamlConfig, error) {
	yamlConfig, err := config.NewYamlConfig("../../config/config.yaml")
	config.YamlConfigGlobal = yamlConfig
	if err != nil {
		return nil, err
	}
  utils.InitLogger()
	return yamlConfig, err
}
