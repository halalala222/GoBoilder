package repository

import _ "embed"

//go:embed gorm.go.tmpl
var GormTemplate []byte

//go:embed mongodb.go.tmpl
var MongodbTemplate []byte

//go:embed mysql.go.tmpl
var MysqlTemplate []byte

//go:embed postgresql.go.tmpl
var PostgresqlTemplate []byte

//go:embed sqlite.go.tmpl
var SqliteTemplate []byte
