package build

import (
	dbConfig "github.com/halalala222/GoBoilder/internal/config"
)

var _ Builder = &ConfigBuilder{}

type ConfigBuilder struct {
	projectName string
	modulePath  string
	db          string
	library     string
}

func (c *ConfigBuilder) String() string {
	return "ConfigBuilder"
}

func NewConfigBuilder(projectName, modulePath, db, library string) *ConfigBuilder {
	return &ConfigBuilder{
		projectName: projectName,
		modulePath:  modulePath,
		db:          db,
		library:     library,
	}
}

func (c *ConfigBuilder) newDBConfigFileBuilder() (*templateFileBuilder, error) {
	var (
		err           error
		dbLibraryInfo *dbConfig.DBLibraryInfo
	)

	if dbLibraryInfo, err = dbConfig.GetDBLibraryInfo(c.db, c.library); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileName: dbLibraryInfo.FileName,
		template: dbLibraryInfo.Template,
	}, nil
}

func (c *ConfigBuilder) getAllConfigFileBuilder() ([]*templateFileBuilder, error) {
	var (
		allConfigFileBuilder = make([]*templateFileBuilder, 0, 3)
		dbConfigFileBuilder  *templateFileBuilder
		err                  error
	)

	if dbConfigFileBuilder, err = c.newDBConfigFileBuilder(); err != nil {
		return nil, err
	}

	allConfigFileBuilder = append(allConfigFileBuilder, dbConfigFileBuilder)

	return allConfigFileBuilder, nil
}

func (c *ConfigBuilder) Build() error {
	var (
		err         error
		fileBuilder []*templateFileBuilder
	)

	if fileBuilder, err = c.getAllConfigFileBuilder(); err != nil {
		return err
	}

	for _, fileBuild := range fileBuilder {
		if err = fileBuild.build(c.projectName); err != nil {
			return err
		}
	}

	return nil
}
