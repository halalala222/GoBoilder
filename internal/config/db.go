package config

var supportedDB = map[string][]string{
	"MySQL":      {"GORM", "go-sql-driver/mysql"},
	"PostgreSQL": {"GORM", "jackc/pgx"},
	"SQLite":     {"GORM", "mattn/go-sqlite3"},
	"MongoDB":    {"mongo-go-driver/mongo"},
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
		libraries []string
		ok        bool
	)

	if libraries, ok = supportedDB[db]; !ok {
		return make([]string, 0)
	}

	return libraries
}
