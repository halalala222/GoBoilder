package config

import "github.com/halalala222/GoBoilder/internal/template/config"

var _ FileTemplate = &YAMLTemplate{}

type YAMLTemplate struct{}

func (y *YAMLTemplate) Build() []byte {
	return config.YAMLTemplate
}
