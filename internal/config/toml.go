package config

import "github.com/halalala222/GoBoilder/internal/template/config"

var _ FileTemplate = &TOMLTemplate{}

type TOMLTemplate struct{}

func (t *TOMLTemplate) Build() []byte {
	return config.TOMLTemplate
}
