package config

import _ "embed"

//go:embed gorm.go.tmpl
var GormConfigTemplate []byte

//go:embed database_sql_mysql.go.tmpl
var DatabaseSQLMySQLConfigTemplate []byte

//go:embed database_sql_postgresql.go.tmpl
var DatabaseSQLPostgreSQLConfigTemplate []byte

//go:embed database_sql_sqlite.go.tmpl
var DatabaseSQLSQLiteConfigTemplate []byte

//go:embed mongo_driver_mongodb.go.tmpl
var MongoDriverMongoDBConfigTemplate []byte

//go:embed gin.go.tmpl
var GinConfigTemplate []byte
