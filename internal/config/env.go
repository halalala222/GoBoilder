package config

import "github.com/halalala222/GoBoilder/internal/template/config"

var _ FileTemplate = &ENVTemplate{}

type ENVTemplate struct{}

func (e *ENVTemplate) Build() []byte {
	return config.ENVTemplate
}
