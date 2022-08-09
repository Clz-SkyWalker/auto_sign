package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIReaderGetSeed(t *testing.T) {
	yamlConfig, err := testInit()
	assert.Equal(t, err, nil, "配置错误")
	ireader := NewIReaderSign(yamlConfig.Ireader)
	for _, item := range ireader.ireaderList {
		item.getSeed()
		assert.Equal(t, err, nil, err)
		assert.NotEqual(t, item.seed, "", "获取种子为空")
	}
}
