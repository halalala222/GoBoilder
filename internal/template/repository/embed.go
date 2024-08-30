package repository

import _ "embed"

//go:embed gorm.go.tmpl
var GormTemplate []byte

//go:embed mongodb.go.tmpl
var MongoDBTemplate []byte

//go:embed mysql.go.tmpl
var MySQLTemplate []byte

//go:embed postgresql.go.tmpl
var PostgreSQLTemplate []byte

//go:embed sqlite.go.tmpl
var SqliteTemplate []byte
