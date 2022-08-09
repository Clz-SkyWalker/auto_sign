package main

import (
	"auto_sign/pkg/config"
	"auto_sign/pkg/utils"
	"auto_sign/src/container"
	"errors"
	"fmt"
)

func main() {
	utils.InitLogger()
	defer utils.DeferSync()
	utils.AddLogger(utils.InfoSystemStart)
	var err error
	config.YamlConfigGlobal, err = config.NewYamlConfig("./config/config.yaml")
	err = errors.New("t")
	if err != nil {
		utils.AddLogger(utils.InfoSystemStartFail)
		return
	}
	con, err := container.NewSignContainer(container.SignContainerParam{
		ServerJKey: config.YamlConfigGlobal.Serverj,
		CookieList: config.YamlConfigGlobal.Juejin,
	})
	if err != nil {
		fmt.Println("[启动失败]")
		fmt.Println("错误：", err)
		return
	}
	fmt.Println("[启动成功]")
	con.Start()
}
