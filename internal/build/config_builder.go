package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config"
)

var _ Builder = &ConfigBuilder{}

type ConfigBuilder struct {
	projectName string
	modulePath  string
	db          string
}

func (c *ConfigBuilder) String() string {
	return "ConfigBuilder"
}

func NewConfigBuilder(projectName, modulePath, db string) *ConfigBuilder {
	return &ConfigBuilder{
		projectName: projectName,
		modulePath:  modulePath,
		db:          db,
	}
}

func newGormConfigFileBuilder(db string, modulePath string) *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.GormConfigFileName,
		template: config.GormConfigTemplate,
		data: &struct {
			ModulePath string
			DB         string
		}{
			ModulePath: modulePath,
			DB:         db,
		},
	}
}

func getAllConfigFileBuilder(db string, modulePath string) []*templateFileBuilder {
	return []*templateFileBuilder{
		newGormConfigFileBuilder(db, modulePath),
	}
}

func (c *ConfigBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range getAllConfigFileBuilder(c.db, c.modulePath) {
		if err = fileBuild.build(filepath.Join(c.projectName, constants.ProjectConfigPkgPath)); err != nil {
			return err
		}
	}

	return nil
}
