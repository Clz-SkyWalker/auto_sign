package container

import (
	"auto_sign/pkg/config"
	"auto_sign/pkg/recyle"
	"auto_sign/src/service"
	"fmt"
	"sync"
)

func NewSignContainer(param SignContainerParam) (*SignContainer, error) {
	recyleList := make([]*recyle.Recyle, 0)
	if len(param.CookieList) == 0 {
		juejin := service.NewJueJinSign(param.CookieList)
		if param.RParam.RType == recyle.RecyleType(0) {
			param.RParam = recyle.RecyleParam{
				RType: recyle.EveryDayRang,
				CTime: "3:00",
				Range: 100,
				Fun: func() {
					juejin.Start()
				},
			}
		}
		recyleList = append(recyleList, recyle.NewRecyle(param.RParam))
	}

	return &SignContainer{
		recyleList: recyleList,
	}, nil
}

type SignContainerParam struct {
	ServerJKey string
	CookieList []string
	RParam     recyle.RecyleParam
}

type SignContainer struct {
	yaml       *config.YamlConfig
	recyleList []*recyle.Recyle
}

func (c *SignContainer) Start() {
	if len(c.recyleList) == 0 {
		fmt.Println("no auto sign")
		return
	}
	for _, item := range c.recyleList {
		item.CreateRecyle()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
