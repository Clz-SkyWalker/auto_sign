package recyle

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RecyleType int

const (
	EveryDayRang RecyleType = iota + 1
)

func NewRecyle(param RecyleParam) *Recyle {
	return &Recyle{
		RecyleParam: param,
	}
}

type RecyleParam struct {
	RType RecyleType // 循环类型
	CTime string     // 时间控制
	Range int        // +分钟范围
	Fun   func()     // 需要执行的方法
}

type Recyle struct {
	RecyleParam
	nextTime time.Time // 下一次执行时间
	duration time.Duration
	tick     *time.Ticker
	Err      error
}

// 创建一个循环执行
func (r *Recyle) CreateRecyle() {
	switch r.RType {
	case EveryDayRang:
		go r.everyDay()
	default:
	}
}

func (r *Recyle) everyDay() {
	r.getNextTime()
	r.tick = time.NewTicker(r.duration)
	for {
		select {
		case <-r.tick.C:
			r.Fun()
			r.getNextTime()
			r.tick.Reset(r.duration)
		}
	}
}

// 获取下一次执行时间
func (r *Recyle) getNextTime() {
	timeSplit := strings.Split(r.CTime, ":")
	if len(timeSplit) != 2 {
		r.Err = errors.New("时间格式错误")
		return
	}
	hour, err := strconv.Atoi(timeSplit[0])
	if err != nil {
		r.Err = err
		return
	}
	minute, err := strconv.Atoi(timeSplit[1])
	if err != nil {
		r.Err = err
		return
	}
	now := time.Now()
	var minuteNum int
	if r.Range > 0 {
		minuteNum = rand.Intn(r.Range)
	}
	switch {
	case now.Hour() > hour || (now.Hour() == hour && now.Minute() >= minute):
		r.nextTime = time.Date(now.Year(), now.Month(), now.Day()+1, hour, minute, 0, 0, time.Local).Add(
			time.Duration(time.Minute * time.Duration(minuteNum)))
		r.duration = r.nextTime.Sub(now)
	default:
		r.nextTime = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local).Add(
			time.Duration(time.Minute * time.Duration(minuteNum)))
		r.duration = r.nextTime.Sub(now)
	}
}
