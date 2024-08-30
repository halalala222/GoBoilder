package repository

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var db2FileInfo = map[string]*template.FileInfo{
	constants.DataBaseMySQL + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
	},
	constants.DataBaseMySQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: mysqlTemplate,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: postgresqlTemplate,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryDatabaseSQL: {
		Template: sqliteTemplate,
	},
	constants.DataBaseMongoDB + constants.DatabaseLibraryMongoDriver: {
		Template: mongoDBTemplate,
	},
}

// GetRepositoryFileTemplateInfo returns the repository file template for the given database and library.
func GetRepositoryFileTemplateInfo(database, dbLibrary string) (*template.FileInfo, error) {
	if fileInfo, ok := db2FileInfo[database+dbLibrary]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrDBOrDBLibraryNotSupported
}
