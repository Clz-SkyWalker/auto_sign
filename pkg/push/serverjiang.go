package push

import (
	"auto_sign/pkg/utils"
	"strings"
)

func NewPushServerj(param PushServerJParam) *PushServerJ {
	return &PushServerJ{
		PushServerJParam: param,
	}
}

type PushServerJParam struct {
	Key   string
	Title string
	Desp  string
}

type PushServerJ struct {
	PushServerJParam
	request *utils.UtilsRequest
}

func (p *PushServerJ) Push() error {
	url := strings.Replace(utils.ServerJBase, utils.Key, p.Key, 1)
	p.request = utils.NewUtilsRequest(utils.UtilsRequestParam{
		Url: url,
		Params: map[string]string{
			utils.Title: p.Title,
			utils.Desp:  p.Desp,
		},
	})

	p.request.Get()
	return nil
}
