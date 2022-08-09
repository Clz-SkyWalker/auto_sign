package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"auto_sign/pkg/config"
	"auto_sign/pkg/push"
	"auto_sign/pkg/utils"
)

func NewJueJinSign(cookieList []string) *JueJinSign {
	itemList := make([]*signItem, 0, len(cookieList))
	for _, item := range cookieList {
		itemList = append(itemList, &signItem{cookie: item})
	}
	return &JueJinSign{
		cookieList: cookieList,
		signList:   itemList,
	}
}

type JueJinSign struct {
	cookieList []string
	signList   []*signItem
	ResultInfo string
	request    utils.UtilsRequest
}

type signItem struct {
	name         string
	isSign       bool
	freeDrwa     float64 // 免费抽奖次数
	cookie       string
	incrPoint    float64 // 签到增加的矿石
	totalPoint   float64 // 总矿石数
	continousDay float64 // 连续签到
	totalSignDay float64 // 累计签到
	gift         string  // 礼品名称
	err          error
}

type juejinResponse struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

type juejinName struct {
	UserName string `json:"user_name"`
}

func (item *signItem) getHeader(url string) http.Header {
	return http.Header{
		"Accept":       []string{"*/*"},
		"Content-type": []string{"application/json"},
		"Referer":      []string{url},
		"Cookie":       []string{item.cookie},
		"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"},
	}
}

func (j *JueJinSign) Start() {
	j.process()
	j.push()
}

func (j *JueJinSign) push() {
	serverj := push.NewPushServerj(push.PushServerJParam{
		Key:   config.YamlConfigGlobal.Serverj,
		Title: "[掘金]",
		Desp:  j.ResultInfo,
	})
	serverj.Push()
}

func (j *JueJinSign) process() {
	builder := strings.Builder{}
	builder.WriteString("[掘金签到开始]\n")
	for _, item := range j.signList {
		cookie := strings.Split(item.cookie, ";")[0]
		builder.WriteString(fmt.Sprintf("[cookie]:%s\n", cookie))
		j.getName(item)
		if item.err != nil {
			builder.WriteString(fmt.Sprintf("[err]:%s\n", item.err))
			continue
		}
		j.checkSignStatus(item)
		if item.err != nil {
			builder.WriteString(fmt.Sprintf("[err]:%s\n", item.err))
			continue
		}
		if item.isSign {
			j.getTotalPoint(item)
		} else {
			j.sign(item)
			if item.err != nil {
				builder.WriteString(fmt.Sprintf("[err]:%s\n", item.err))
				continue
			}
		}

		j.checkLuckyDraw(item)
		if item.err != nil {
			builder.WriteString(fmt.Sprintf("[err]:%s\n", item.err))
			continue
		}
		if item.freeDrwa > 0 {
			j.luckyDraw(item)
			if item.err != nil {
				builder.WriteString(fmt.Sprintf("[err]:%s\n", item.err))
				continue
			}
		}
		j.getTotalSignDay(item)

		builder.WriteString(fmt.Sprintf(`[名字]:%s
      [签到状态]:%s
      [总矿石]:%.0f
      [签到获取矿石]:%.0f
      [抽奖礼品]:%s
      [总签到天数]:%.0f
      [连续签到天数]:%.0f
      `, item.name, strconv.FormatBool(item.isSign), item.totalPoint, item.incrPoint, item.gift, item.totalSignDay, item.continousDay))
		builder.WriteString("\n")
	}

	builder.WriteString("[掘金签到结束]\n")
	j.ResultInfo = builder.String()
}

// 获取名字
func (j *JueJinSign) getName(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_GetName,
		Header: item.getHeader(utils.JUEJIN_GetName),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Get()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}
	value, ok := resp.Data.(map[string]interface{})["user_name"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
		return
	}
	item.name = value.(string)
}

// 检测签到状态
func (j *JueJinSign) checkSignStatus(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_CheckSign,
		Header: item.getHeader(utils.JUEJIN_CheckSign),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Get()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}
	item.isSign = resp.Data.(bool)
}

// 签到
func (j *JueJinSign) sign(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_Sign,
		Header: item.getHeader(utils.JUEJIN_Sign),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Post()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}

	incrPoint, ok := resp.Data.(map[string]interface{})["incr_point"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.incrPoint = incrPoint.(float64)
	sumPoint, ok := resp.Data.(map[string]interface{})["sum_point"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.totalPoint = sumPoint.(float64)
}

// 检测免费抽奖次数
func (j *JueJinSign) checkLuckyDraw(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_CheckFreeDraw,
		Header: item.getHeader(utils.JUEJIN_CheckFreeDraw),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Get()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}
	value, ok := resp.Data.(map[string]interface{})["free_count"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.freeDrwa, ok = value.(float64)
}

// 抽奖
/**
 * 抽奖函数
 * 目前已知奖品
 * lottery_id: 6981716980386496552、name: 66矿石、type: 1
 * lottery_id: 6981716405976743943、name: Bug、type: 2
 * lottery_id: 7020245697131708419、name: 掘金帆布袋、type: 4
 * lottery_id: 7017679355841085472、name: 随机限量徽章、type: 4
 * lottery_id: 6997270183769276416、name: Yoyo抱枕、type: 4
 * lottery_id: 7001028932350771203、name: 掘金马克杯、type: 4
 * lottery_id: 7020306802570952718、name: 掘金棒球帽、type: 4
 * lottery_id: 6981705951946489886、name: Switch、type: 3
 */
func (j *JueJinSign) luckyDraw(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_Draw,
		Header: item.getHeader(utils.JUEJIN_Draw),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Post()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}

	gift, ok := resp.Data.(map[string]interface{})["lottery_name"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.gift = gift.(string)
}

// 获取矿石总数
func (j *JueJinSign) getTotalPoint(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_Total,
		Header: item.getHeader(utils.JUEJIN_Total),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Get()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}
	item.totalPoint = resp.Data.(float64)
}

// 获取签到天数
func (j *JueJinSign) getTotalSignDay(item *signItem) {
	param := utils.UtilsRequestParam{
		Url:    utils.JUEJIN_TotalSignDay,
		Header: item.getHeader(utils.JUEJIN_TotalSignDay),
	}

	j.request = *utils.NewUtilsRequest(param)
	j.request.Get()
	if j.request.Err.Code != 0 {
		item.err = j.request.Err
		return
	}
	byteList, err := ioutil.ReadAll(j.request.Respose.Body)
	if err != nil {
		item.err = err
		return
	}
	var resp juejinResponse
	err = json.Unmarshal(byteList, &resp)
	if err != nil {
		item.err = err
		return
	}
	if resp.ErrNo != 0 {
		item.err = fmt.Errorf("code:%d,msg:%s", resp.ErrNo, resp.ErrMsg)
		return
	}
	contCount, ok := resp.Data.(map[string]interface{})["cont_count"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.continousDay = contCount.(float64)

	sumCount, ok := resp.Data.(map[string]interface{})["sum_count"]
	if !ok {
		item.err = fmt.Errorf("%s%s", "响应结构改变：", utils.JUEJIN_CheckFreeDraw)
	}
	item.totalSignDay = sumCount.(float64)
}
