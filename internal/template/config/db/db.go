package db

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var supportedDB2DBLibrary = map[string][]string{
	constants.DataBaseMySQL:      {constants.DatabaseLibraryGorm, constants.DatabaseLibraryDatabaseSQL},
	constants.DataBasePostgreSQL: {constants.DatabaseLibraryGorm, constants.DatabaseLibraryDatabaseSQL},
	constants.DataBaseSQLite:     {constants.DatabaseLibraryGorm, constants.DatabaseLibraryDatabaseSQL},
	constants.DataBaseMongoDB:    {constants.DatabaseLibraryMongoDriver},
}

var supportedDBLibrary2FileInfo = map[string]*template.FileInfo{
	constants.DataBaseMySQL + constants.DatabaseLibraryGorm: {
		Template: GormConfigTemplate,
		FileName: constants.GormConfigFileName,
	},
	constants.DataBaseMySQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: DatabaseSQLMySQLConfigTemplate,
		FileName: constants.DatabaseSQLMySQLConfigFileName,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryGorm: {
		Template: GormConfigTemplate,
		FileName: constants.GormConfigFileName,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: DatabaseSQLPostgreSQLConfigTemplate,
		FileName: constants.DatabaseSQLPostgreSQLConfigFileName,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryGorm: {
		Template: GormConfigTemplate,
		FileName: constants.GormConfigFileName,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryDatabaseSQL: {
		Template: DatabaseSQLSQLiteConfigTemplate,
		FileName: constants.DatabaseSQLSQLiteConfigFileName,
	},
	constants.DataBaseMongoDB + constants.DatabaseLibraryMongoDriver: {
		Template: MongoDriverMongoDBConfigTemplate,
		FileName: constants.MongoDriverMongoDBConfigFileName,
	},
}

// GetDBLibraries returns the supported libraries for the given database.
func GetDBLibraries(database string) []string {
	return supportedDB2DBLibrary[database]
}

// GetSupportedDatabases returns the supported databases.
func GetSupportedDatabases() []string {
	var databases = make([]string, 0, len(supportedDB2DBLibrary))

	for database := range supportedDB2DBLibrary {
		databases = append(databases, database)
	}

	return databases
}

// GetDBLibraryFileTemplateInfo returns the FileInfo for the db library file template.
func GetDBLibraryFileTemplateInfo(database, library string) (*template.FileInfo, error) {
	if fileInfo, ok := supportedDBLibrary2FileInfo[database+library]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrDBLibraryNotSupported
}
