package repository

import "github.com/halalala222/GoBoilder/internal/constants"

type Template interface {
	Build() []byte
}

func db2Template() map[string]Template {
	return map[string]Template{
		constants.DataBaseMySQL + constants.DatabaseLibraryGorm:             &GormRepositoryTemplate{},
		constants.DataBaseMySQL + constants.DatabaseLibraryDatabaseSQL:      &MySQLRepositoryTemplate{},
		constants.DataBasePostgreSQL + constants.DatabaseLibraryGorm:        &GormRepositoryTemplate{},
		constants.DataBasePostgreSQL + constants.DatabaseLibraryDatabaseSQL: &PostgreSQLRepositoryTemplate{},
		constants.DataBaseSQLite + constants.DatabaseLibraryGorm:            &GormRepositoryTemplate{},
		constants.DataBaseSQLite + constants.DatabaseLibraryDatabaseSQL:     &SQLiteRepositoryTemplate{},
		constants.DataBaseMongoDB + constants.DatabaseLibraryMongoDriver:    &MongoDBRepositoryTemplate{},
	}
}

// GetRepositoryTemplate returns the repository template based on the database
func GetRepositoryTemplate(database string, dbLibrary string) (Template, error) {
	var (
		dbTemplate = db2Template()
		template   Template
		ok         bool
	)

	if template, ok = dbTemplate[database+dbLibrary]; !ok {
		return nil, constants.ErrDBOrDBLibraryNotSupported
	}

	return template, nil
}
