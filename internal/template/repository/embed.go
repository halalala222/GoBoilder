package repository

import _ "embed"

//go:embed gorm.go.tmpl
var GormTemplate []byte
