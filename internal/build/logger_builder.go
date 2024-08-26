package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/logger"
)

var _ Builder = &LoggerBuilder{}

type LoggerBuilder struct {
	projectName string
	library     string
	modulePath  string
}

func (l *LoggerBuilder) String() string {
	return "LoggerBuilder"
}

func NewLoggerBuilder(projectName, library, modulePath string) *LoggerBuilder {
	return &LoggerBuilder{
		projectName: projectName,
		library:     library,
		modulePath:  modulePath,
	}
}

func newLoggerFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.LoggerFileNae,
		template: logger.TemplateLogger(),
	}
}

func newLoggerLibraryFileBuilder(library string) *templateFileBuilder {
	var (
		libraryTemplate logger.LibraryTemplate
		err             error
	)

	if libraryTemplate, err = logger.GetLibraryTemplate(library); err != nil {
		return nil
	}

	return &templateFileBuilder{
		fileName: logger.GetLoggerLibraryFileName(library),
		template: libraryTemplate.Build(),
	}
}

func newLogFileBuilder(modulePath string) *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.LogFileName,
		template: logger.TemplateLog(),
		data: &struct {
			ModulePath string
		}{
			ModulePath: modulePath,
		},
	}
}

func getAllPkgLoggerFileBuilder(library string) []*templateFileBuilder {
	return []*templateFileBuilder{
		newLoggerFileBuilder(),
		newLoggerLibraryFileBuilder(library),
	}
}

func getAllInternalPkgLoggerFileBuilder(modulePath string) []*templateFileBuilder {
	return []*templateFileBuilder{
		newLogFileBuilder(modulePath),
	}
}

func (l *LoggerBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range getAllPkgLoggerFileBuilder(l.library) {
		if err = fileBuild.build(filepath.Join(l.projectName, constants.ProjectLoggerPkgPath)); err != nil {
			return err
		}
	}

	for _, fileBuild := range getAllInternalPkgLoggerFileBuilder(l.modulePath) {
		if err = fileBuild.build(filepath.Join(l.projectName, constants.ProjectInternalPkgLogPath)); err != nil {
			return err
		}
	}

	return nil
}
