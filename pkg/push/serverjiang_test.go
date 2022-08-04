package push

import (
	"testing"
)

func TestServerjPush(t *testing.T) {
	serverj := NewPushServerj(PushServerJParam{Key: "SCT54442TB9U33jkBFcb8qqN6lnoaTq4U",
		Title: "test", Desp: "我是一个测试"})
  serverj.Push()
}
