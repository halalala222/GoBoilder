package constants

const (
	ZapLoggerLibrary  = "zap"
	SlogLoggerLibrary = "slog"

	NoneCurrentBuildInfo = "(none)"
	CurrentBuildTitle    = "Current build"

	ApplicationHeader = "GoBoiler Build"

	LoggerKey               = "logger"
	ChoiceLoggerTitle       = "Choose your logger"
	ChoiceLoggerDescription = "This will determine your logger"

	HTTPFrameKey               = "http_frame"
	ChoiceHTTPFrameTitle       = "Choose your http frame"
	ChoiceHTTPFrameDescription = "This will determine your http frame"

	DoneKey = "done"
	AllDone = "All done?"

	Yep  = "Yep"
	Quit = "Quit"

	QuitBody = "Application quit"

	DBKey               = "db"
	ChoiceDBTitle       = "Choose your database"
	ChoiceDBDescription = "This will determine your database"

	DBLibraryKey               = "db_library"
	ChoiceDBLibraryTitle       = "Choose your database library"
	ChoiceDBLibraryDescription = "This will determine your database library"

	ProjectNameKey         = "project_name"
	ProjectName            = "Enter project name"
	ProjectNamePlaceholder = "my_project"
	ProjectNameDescription = "This will be the name of your project"

	ModulePathPrefixKey         = "module_path"
	ModulePathPrefix            = "Enter module path"
	ModulePathPrefixPlaceholder = "github.com/halalala222"
	ModulePathPrefixDescription = "module path prefix + project name = module path"

	ProjectLoggerPkgPath      = "/pkg/log"
	ProjectInternalPkgLogPath = "/internal/log"
	ProjectDomainPkgPath      = "/domain"
	ProjectConfigPkgPath      = "/internal/config"
	ProjectRepositoryPkgPath  = "/internal/repository"
	ProjectUserServicePkgPath = "/user"

	ConfigFileTypeKey               = "config_file_type"
	ChoiceConfigFileTypeTitle       = "Choose your config file type"
	ChoiceConfigFileTypeDescription = "This will determine your config file type"

	GitIgnoreFileName                   = ".gitignore"
	READEMEFileName                     = "README.md"
	MakefileFileName                    = "Makefile"
	LoggerFileNae                       = "logger.go"
	SlogLoggerFileName                  = "slog_logger.go"
	ZapLoggerFileName                   = "zap_logger.go"
	LogFileName                         = "log.go"
	DomainUserFileNae                   = "user.go"
	DomainErrorsFileName                = "errors.go"
	GormConfigFileName                  = "gorm.go"
	DatabaseSQLMySQLConfigFileName      = "mysql.go"
	DatabaseSQLPostgreSQLConfigFileName = "postgresql.go"
	DatabaseSQLSQLiteConfigFileName     = "sqlite.go"
	MongoDriverMongoDBConfigFileName    = "mongodb.go"
	RepositoryFileName                  = "user.go"
	ServiceFileName                     = "service.go"

	DatabaseLibraryGorm        = "gorm"
	DatabaseLibraryMongoDriver = "mongo-driver"
	DatabaseLibraryDatabaseSQL = "database/sql"

	DataBaseMySQL      = "MySQL"
	DataBasePostgreSQL = "PostgreSQL"
	DataBaseSQLite     = "SQLite"
	DataBaseMongoDB    = "MongoDB"

	ConfigFileName     = "config.go"
	YAMLConfigFileType = "yaml"
	YAMLConfigFileName = "config.yaml"
	JSONConfigFileType = "json"
	JSONConfigFileName = "config.json"
	TOMLConfigFileType = "toml"
	TOMLConfigFileName = "config.toml"
	ENVConfigFileType  = "env"
	ENVConfigFileName  = ".env"
)
