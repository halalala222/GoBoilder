package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/repository"
)

var db2FileInfo = map[string]*FileInfo{
	constants.DataBaseMySQL + constants.DatabaseLibraryGorm: {
		Template: repository.GormTemplate,
	},
	constants.DataBaseMySQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: repository.MySQLTemplate,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryGorm: {
		Template: repository.GormTemplate,
	},
	constants.DataBasePostgreSQL + constants.DatabaseLibraryDatabaseSQL: {
		Template: repository.PostgreSQLTemplate,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryGorm: {
		Template: repository.GormTemplate,
	},
	constants.DataBaseSQLite + constants.DatabaseLibraryDatabaseSQL: {
		Template: repository.SqliteTemplate,
	},
	constants.DataBaseMongoDB + constants.DatabaseLibraryMongoDriver: {
		Template: repository.MongoDBTemplate,
	},
}

// GetRepositoryFileTemplateInfo returns the repository file template for the given database and library.
func GetRepositoryFileTemplateInfo(database, dbLibrary string) (*FileInfo, error) {
	if fileInfo, ok := db2FileInfo[database+dbLibrary]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrDBOrDBLibraryNotSupported
}
