package build

import (
	"path/filepath"
	"strings"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
	"github.com/halalala222/GoBoilder/internal/template/cmd"
)

var _ Builder = &AppBuilder{}

type AppBuilder struct {
	projectName   string
	modulePath    string
	httpFramework string
	db            string
	dbLibrary     string
}

func (a *AppBuilder) String() string {
	return "AppBuilder"
}

func NewAppBuilder(projectName, modulePath, httpFramework, db, dbLibrary string) *AppBuilder {
	return &AppBuilder{
		projectName:   projectName,
		modulePath:    modulePath,
		httpFramework: httpFramework,
		db:            db,
		dbLibrary:     dbLibrary,
	}
}

func (a *AppBuilder) appTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(a.projectName, constants.ProjectCmdPkgPath),
		Data: struct {
			ModulePath    string
			HTTPFramework string
			DB            string
			DBLibrary     string
		}{
			ModulePath:    a.modulePath,
			HTTPFramework: a.httpFramework,
			DB:            strings.ToLower(a.db),
			DBLibrary:     a.dbLibrary,
		},
	}
}

func (a *AppBuilder) newAppFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  cmd.GetAppGoFileInfo(),
		buildInfo: a.appTemplateBuildInfo(),
	}
}

func (a *AppBuilder) Build() error {
	var (
		appFileBuilder = a.newAppFileBuilder()
	)

	return appFileBuilder.fileInfo.Build(appFileBuilder.buildInfo)
}
