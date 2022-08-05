package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestYamlParse(t *testing.T) {
	_, err := NewYamlConfig("../../config/config.yaml")
	assert.NotEqual(t, err, nil)
}
