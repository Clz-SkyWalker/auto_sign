package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"auto_sign/pkg/utils"
)

func NewIReaderSign(uid []string) *IReaderSign {
	ireaderList := make([]ireaderItem, 0, len(uid))
	for _, item := range uid {
		ireaderList = append(ireaderList, ireaderItem{uid: item})
	}
	return &IReaderSign{
		ireaderList: ireaderList,
	}
}

type IReaderSign struct {
	ireaderList []ireaderItem
	ResultInfo  string
}

type ireaderItem struct {
	uid        string // 用户 id
	resultInfo string
	seed       string // 种子
	seedId     string // 真正的id
	err        utils.Errno
	request    *utils.UtilsRequest
}

type ireaderResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (item *ireaderItem) getHeader(url string) http.Header {
	head := http.Header{}
	head.Add("Content-type", "application/json;charset=utf-8")
	head.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.47")
	return head
}

// 获取种子
func (r *ireaderItem) getSeed() {
	r.request = utils.NewUtilsRequest(utils.UtilsRequestParam{
		Url:    utils.IREADER_Seed + r.uid,
		Header: r.getHeader(utils.IREADER_Seed + r.uid),
	})
	r.request.Get()
	if r.request.Err.Code != 0 {
		r.err = r.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(r.request.Respose.Body)
	if err != nil {
		r.err = utils.ErrJsonMarshal.WithErr(err)
		utils.AddLogger(r.err)
		return
	}
	var resp ireaderResp
	fmt.Println(string(byteList))
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		r.err = utils.ErrJsonUnmarshal.WithErr(err)
		utils.AddLogger(r.err)
		return
	}
	if resp.Code != 0 {
		r.resultInfo += fmt.Sprintf("[seed err]\n[code]:%d\n[msg]:%s\n", resp.Code, resp.Msg)
	}
	seed, ok := resp.Data.(map[string]interface{})["seed"].(string)
	if !ok {
		r.err = utils.ErrIReaderSeed
		utils.AddLogger(r.err)
		return
	}
	r.seed = seed
}

// 获取种子
func (r *ireaderItem) sign() {
	r.request = utils.NewUtilsRequest(utils.UtilsRequestParam{
		Url:    utils.IREADER_Sign+ r.uid,
		Header: r.getHeader(utils.IREADER_Seed + r.uid),
	})
	r.request.Get()
	if r.request.Err.Code != 0 {
		r.err = r.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(r.request.Respose.Body)
	if err != nil {
		r.err = utils.ErrJsonMarshal.WithErr(err)
		utils.AddLogger(r.err)
		return
	}
	var resp ireaderResp
	fmt.Println(string(byteList))
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		r.err = utils.ErrJsonUnmarshal.WithErr(err)
		utils.AddLogger(r.err)
		return
	}
	if resp.Code != 0 {
		r.resultInfo += fmt.Sprintf("[seed err]\n[code]:%d\n[msg]:%s\n", resp.Code, resp.Msg)
	}
	seed, ok := resp.Data.(map[string]interface{})["seed"].(string)
	if !ok {
		r.err = utils.ErrIReaderSeed
		utils.AddLogger(r.err)
		return
	}
	r.seed = seed
}

func (r *ireaderItem) splitSeed() {
	seedList := strings.Split(r.seed, "|")
	r.seedId = seedList[len(seedList)-1]
}
