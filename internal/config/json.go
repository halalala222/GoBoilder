package config

import "github.com/halalala222/GoBoilder/internal/template/config"

var _ FileTemplate = &JSONTemplate{}

type JSONTemplate struct{}

func (j *JSONTemplate) Build() []byte {
	return config.JSONTemplate
}
