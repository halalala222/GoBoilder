package config

import _ "embed"

//go:embed gorm.go.tmpl
var GormConfigTemplate []byte
