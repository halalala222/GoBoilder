package repository

import _ "embed"

//go:embed gorm.go.tmpl
var gormTemplate []byte

//go:embed mongodb.go.tmpl
var mongoDBTemplate []byte

//go:embed mysql.go.tmpl
var mysqlTemplate []byte

//go:embed postgresql.go.tmpl
var postgresqlTemplate []byte

//go:embed sqlite.go.tmpl
var sqliteTemplate []byte
