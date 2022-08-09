package service

import (
	"auto_sign/pkg/push"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJueJinStart(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	result.Start()
	assert.Equal(t, result.request.Err, nil, "请求错误")
}

func TestJueJinSignPush(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	result.process()
	serverj := push.NewPushServerj(push.PushServerJParam{Key: config.Serverj,
		Title: "掘金", Desp: result.ResultInfo})
	serverj.Push()
}

func TestJueJinProcess(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	result.process()
	assert.NotEqual(t, result.ResultInfo, "", "获取信息为空")
}

// 测试获取名字
func TestJueJinGetName(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.getName(item)
		assert.Equal(t, item.err, nil, "获取名字错误")
	}
}

// 测试检测签到状态
func TestJueJinCheckSign(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.checkSignStatus(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 检测免费抽奖次数
func TestJueJinCheckFreeLuckyDraw(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.checkLuckyDraw(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 获取矿石总数
func TestJueJinGetPoint(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.getTotalPoint(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 获取签到天数
func TestJueJinSignDay(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.getTotalSignDay(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 签到
func TestJueJinSign(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.sign(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 测试抽奖
func TestJuejinLuckyDraw(t *testing.T) {
	config, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	result := NewJueJinSign(config.Juejin)
	for _, item := range result.signList {
		result.luckyDraw(item)
		if result.request.Err.Code != 0 {
			fmt.Println(result.request.Err)
			return
		}
	}
}
