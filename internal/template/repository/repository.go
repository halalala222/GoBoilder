package repository

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var db2FileInfo = map[string]*template.FileInfo{
	constants.DataBaseMySQL + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBaseMySQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: mysqlTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: postgresqlTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryGorm: {
		Template: gormTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryDatabaseSQL: {
		Template: sqliteTemplate,
		FileName: constants.RepositoryFileName,
	},
	constants.DataBaseMongoDB + constants.DatabaseLibraryMongoDriver: {
		Template: mongoDBTemplate,
		FileName: constants.RepositoryFileName,
	},
}

// GetRepositoryFileTemplateInfo returns the repository file template for the given database and library.
func GetRepositoryFileTemplateInfo(database, dbLibrary string) (*template.FileInfo, error) {
	if fileInfo, ok := db2FileInfo[database+dbLibrary]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrDBOrDBLibraryNotSupported
}
