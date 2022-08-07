package main

import (
	"auto_sign/pkg/config"
	"auto_sign/src/container"
	"fmt"
)

func main() {
	var err error
	config.YamlConfigGlobal, err = config.NewYamlConfig("./config/config.yaml")
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
