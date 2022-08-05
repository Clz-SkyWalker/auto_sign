package container

import (
	"auto_sign/pkg/config"
	"auto_sign/pkg/recyle"
	"auto_sign/src/service"
	"errors"
	"sync"
)

func NewSignContainer(param SignContainerParam) (*SignContainer, error) {
	if len(param.CookieList) == 0 {
		return nil, errors.New("cookie length equal 0")
	}
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
	return &SignContainer{
		juejin: juejin,
		recyle: recyle.NewRecyle(param.RParam),
	}, nil
}

type SignContainerParam struct {
	ServerJKey string
	CookieList []string
	RParam     recyle.RecyleParam
}

type SignContainer struct {
	yaml   *config.YamlConfig
	recyle *recyle.Recyle
	juejin *service.JueJinSign
}

func (c *SignContainer) Start() {
	c.recyle.CreateRecyle()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
