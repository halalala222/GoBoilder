package build

import (
	"path/filepath"
	"strings"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
	"github.com/halalala222/GoBoilder/internal/template/config"
	"github.com/halalala222/GoBoilder/internal/template/config/db"
	"github.com/halalala222/GoBoilder/internal/template/config/http"
)

var _ Builder = &ConfigBuilder{}

type ConfigBuilder struct {
	projectName    string
	modulePath     string
	db             string
	library        string
	loggerLibrary  string
	configFileType string
	httpFramework  string
}

func (c *ConfigBuilder) String() string {
	return "ConfigBuilder"
}

func NewConfigBuilder(projectName, modulePath, db, library, configFileType, httpFramework, loggerLibrary string) *ConfigBuilder {
	return &ConfigBuilder{
		projectName:    projectName,
		modulePath:     modulePath,
		db:             db,
		library:        library,
		configFileType: configFileType,
		httpFramework:  httpFramework,
		loggerLibrary:  loggerLibrary,
	}
}

func (c *ConfigBuilder) dbConfigFileBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(c.projectName, constants.ProjectConfigPkgPath),
		Data: &struct {
			ModulePath string
			DB         string
		}{
			ModulePath: c.modulePath,
			DB:         strings.ToLower(c.db),
		},
	}
}

func (c *ConfigBuilder) configLoaderFileBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(c.projectName, constants.ProjectConfigPkgPath),
		Data: &struct {
			ModulePath     string
			ConfigFileType string
			ProjectName    string
			DB             string
			LoggerLibrary  string
		}{
			ModulePath:     c.modulePath,
			ConfigFileType: c.configFileType,
			ProjectName:    c.projectName,
			DB:             c.db,
			LoggerLibrary:  c.loggerLibrary,
		},
	}
}

func (c *ConfigBuilder) httpFrameFileBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(c.projectName, constants.ProjectConfigPkgPath),
		Data: &struct {
			ModulePath string
		}{
			ModulePath: c.modulePath,
		},
	}
}

func (c *ConfigBuilder) newDBConfigFileBuilder() (*templateFileBuilder, error) {
	var (
		err      error
		fileInfo *template.FileInfo
	)

	if fileInfo, err = db.GetDBLibraryFileTemplateInfo(c.db, c.library); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: c.dbConfigFileBuildInfo(),
	}, nil
}

func (c *ConfigBuilder) newConfigLoaderFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  config.GetConfigLoaderFileTemplateInfo(),
		buildInfo: c.configLoaderFileBuildInfo(),
	}
}

func (c *ConfigBuilder) getConfigFileBuilder() (*templateFileBuilder, error) {
	var (
		err      error
		fileInfo *template.FileInfo
	)

	if fileInfo, err = config.GetConfigFileTemplateInfo(c.configFileType); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: c.configLoaderFileBuildInfo(),
	}, nil
}

func (c *ConfigBuilder) getHTTPFrameFileBuilder() (*templateFileBuilder, error) {
	var (
		err      error
		fileInfo *template.FileInfo
	)

	if fileInfo, err = http.GetHTTPFrameFileTemplateInfo(c.httpFramework); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: c.httpFrameFileBuildInfo(),
	}, nil
}

func (c *ConfigBuilder) getAllConfigFileBuilder() ([]*templateFileBuilder, error) {
	var (
		allConfigFileBuilder = make([]*templateFileBuilder, 0, 3)
		dbConfigFileBuilder  *templateFileBuilder
		httpFrameFileBuilder *templateFileBuilder
		configLoaderBuilder  = c.newConfigLoaderFileBuilder()
		configFileBuilder    *templateFileBuilder
		err                  error
	)

	if dbConfigFileBuilder, err = c.newDBConfigFileBuilder(); err != nil {
		return nil, err
	}

	if httpFrameFileBuilder, err = c.getHTTPFrameFileBuilder(); err != nil {
		return nil, err
	}

	if configFileBuilder, err = c.getConfigFileBuilder(); err != nil {
		return nil, err
	}

	allConfigFileBuilder = append(allConfigFileBuilder, dbConfigFileBuilder)
	allConfigFileBuilder = append(allConfigFileBuilder, httpFrameFileBuilder)
	allConfigFileBuilder = append(allConfigFileBuilder, configLoaderBuilder)
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
		if err = fileBuild.fileInfo.Build(fileBuild.buildInfo); err != nil {
			return err
		}
	}

	return nil
}
