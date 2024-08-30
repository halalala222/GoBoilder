package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/config"
	"github.com/halalala222/GoBoilder/internal/constants"
)

var _ Builder = &ConfigBuilder{}

type ConfigBuilder struct {
	projectName    string
	modulePath     string
	db             string
	library        string
	configFileType string
	httpFramework  string
}

func (c *ConfigBuilder) String() string {
	return "ConfigBuilder"
}

func NewConfigBuilder(projectName, modulePath, db, library, configFileType, httpFramework string) *ConfigBuilder {
	return &ConfigBuilder{
		projectName:    projectName,
		modulePath:     modulePath,
		db:             db,
		library:        library,
		configFileType: configFileType,
		httpFramework:  httpFramework,
	}
}

func (c *ConfigBuilder) newDBConfigFileBuilder() (*templateFileBuilder, error) {
	var (
		err           error
		dbLibraryInfo *config.DBLibraryInfo
	)

	if dbLibraryInfo, err = config.GetDBLibraryInfo(c.db, c.library); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileName: dbLibraryInfo.FileName,
		template: dbLibraryInfo.Template,
		data: &struct {
			ModulePath string
			DB         string
		}{
			ModulePath: c.modulePath,
			DB:         c.db,
		},
	}, nil
}

func (c *ConfigBuilder) newConfigBuilder() *templateFileBuilder {

	return &templateFileBuilder{
		fileName: constants.ConfigFileName,
		template: config.GetConfigTemplate(),
		data: &struct {
			ConfigFileType string
			ProjectName    string
		}{
			ConfigFileType: c.configFileType,
			ProjectName:    c.projectName,
		},
	}
}

func (c *ConfigBuilder) getConfigFileBuilder() (*templateFileBuilder, error) {
	var (
		err                error
		configFileTemplate config.FileTemplate
		configFileName     string
	)

	if configFileTemplate, err = config.GetConfigFileTemplate(c.configFileType); err != nil {
		return nil, err
	}

	if configFileName, err = config.GetConfigFileName(c.configFileType); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileName: configFileName,
		template: configFileTemplate.Build(),
		data: &struct {
			DB string
		}{
			DB: c.db,
		},
	}, nil
}

func (c *ConfigBuilder) getHTTPFrameFilerBuilder() (*templateFileBuilder, error) {
	var (
		err           error
		httpFrameInfo *config.FrameworkInfo
	)

	if httpFrameInfo, err = config.GetFrameworkInfo(c.httpFramework); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileName: httpFrameInfo.FileName,
		template: httpFrameInfo.Template,
		data: &struct {
			ModulePath string
		}{
			ModulePath: c.modulePath,
		},
	}, nil
}

func (c *ConfigBuilder) getAllConfigFileBuilder() ([]*templateFileBuilder, error) {
	var (
		allConfigFileBuilder = make([]*templateFileBuilder, 0, 3)
		dbConfigFileBuilder  *templateFileBuilder
		httpFrameFileBuilder *templateFileBuilder
		configBuilder        = c.newConfigBuilder()
		configFileBuilder    *templateFileBuilder
		err                  error
	)

	if dbConfigFileBuilder, err = c.newDBConfigFileBuilder(); err != nil {
		return nil, err
	}

	if httpFrameFileBuilder, err = c.getHTTPFrameFilerBuilder(); err != nil {
		return nil, err
	}

	if configFileBuilder, err = c.getConfigFileBuilder(); err != nil {
		return nil, err
	}

	allConfigFileBuilder = append(allConfigFileBuilder, dbConfigFileBuilder)
	allConfigFileBuilder = append(allConfigFileBuilder, httpFrameFileBuilder)
	allConfigFileBuilder = append(allConfigFileBuilder, configBuilder)
	allConfigFileBuilder = append(allConfigFileBuilder, configFileBuilder)

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
		if err = fileBuild.build(filepath.Join(c.projectName, constants.ProjectConfigPkgPath)); err != nil {
			return err
		}
	}

	return nil
}
