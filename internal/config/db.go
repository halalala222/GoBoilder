package config

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config"
)

type DBLibraryInfo struct {
	name     string
	FileName string
	Template []byte
}

var supportedDB = map[string][]*DBLibraryInfo{
	constants.DataBaseMySQL: {
		{
			name:     constants.DatabaseLibraryGorm,
			FileName: constants.GormConfigFileName,
			Template: config.GormConfigTemplate,
		},
		{
			name:     constants.DatabaseLibraryDatabaseSQL,
			FileName: constants.DatabaseSQLMySQLConfigFileName,
			Template: config.DatabaseSQLMySQLConfigTemplate,
		},
	},
	constants.DataBasePostgreSQL: {
		{
			name:     constants.DatabaseLibraryGorm,
			FileName: constants.GormConfigFileName,
			Template: config.GormConfigTemplate,
		},
		{
			name:     constants.DatabaseLibraryDatabaseSQL,
			FileName: constants.DatabaseSQLPostgreSQLConfigFileName,
			Template: config.DatabaseSQLPostgreSQLConfigTemplate,
		},
	},
	constants.DataBaseSQLite: {
		{
			name:     constants.DatabaseLibraryGorm,
			FileName: constants.GormConfigFileName,
			Template: config.GormConfigTemplate,
		},
		{
			name:     constants.DatabaseLibraryDatabaseSQL,
			FileName: constants.DatabaseSQLSQLiteConfigFileName,
			Template: config.DatabaseSQLSQLiteConfigTemplate,
		},
	},
	constants.DataBaseMongoDB: {
		{
			name:     constants.DatabaseLibraryMongoDriver,
			FileName: constants.MongoDriverMongoDBConfigFileName,
			Template: config.MongoDriverMongoDBConfigTemplate,
		},
	},
}

// GetSupportedDB returns the supported databases.
func GetSupportedDB() []string {
	var dbs = make([]string, 0, len(supportedDB))

	for db := range supportedDB {
		dbs = append(dbs, db)
	}

	return dbs
}

// GetDBLibraries returns the libraries for the given database.
func GetDBLibraries(db string) []string {
	var (
		libraries   []string
		ok          bool
		dbLibraries []*DBLibraryInfo
	)

	if dbLibraries, ok = supportedDB[db]; !ok {
		return make([]string, 0)
	}

	for _, library := range dbLibraries {
		libraries = append(libraries, library.name)
	}

	return libraries
}

// GetDBLibraryInfo returns the library info for the given database and library.
func GetDBLibraryInfo(db, library string) (*DBLibraryInfo, error) {
	var (
		dbLibraries []*DBLibraryInfo
		ok          bool
	)

	if dbLibraries, ok = supportedDB[db]; !ok {
		return nil, constants.ErrDBNotSupported
	}

	for _, lib := range dbLibraries {
		if lib.name == library {
			return lib, nil
		}
	}

	return nil, constants.ErrDBLibraryNotSupported
}
