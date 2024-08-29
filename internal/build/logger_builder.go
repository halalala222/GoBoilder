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

func (l *LoggerBuilder) newLoggerFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.LoggerFileNae,
		template: logger.TemplateLogger(),
	}
}

func (l *LoggerBuilder) newLoggerLibraryFileBuilder() *templateFileBuilder {
	var (
		libraryTemplate logger.LibraryTemplate
		err             error
	)

	if libraryTemplate, err = logger.GetLibraryTemplate(l.library); err != nil {
		return nil
	}

	return &templateFileBuilder{
		fileName: logger.GetLoggerLibraryFileName(l.library),
		template: libraryTemplate.Build(),
	}
}

func (l *LoggerBuilder) newLogFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.LogFileName,
		template: logger.TemplateLog(),
		data: &struct {
			ModulePath string
		}{
			ModulePath: l.modulePath,
		},
	}
}

func (l *LoggerBuilder) getAllPkgLoggerFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		l.newLoggerFileBuilder(),
		l.newLoggerLibraryFileBuilder(),
	}
}

func (l *LoggerBuilder) getAllInternalPkgLoggerFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		l.newLogFileBuilder(),
	}
}

func (l *LoggerBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range l.getAllPkgLoggerFileBuilder() {
		if err = fileBuild.build(filepath.Join(l.projectName, constants.ProjectLoggerPkgPath)); err != nil {
			return err
		}
	}

	for _, fileBuild := range l.getAllInternalPkgLoggerFileBuilder() {
		if err = fileBuild.build(filepath.Join(l.projectName, constants.ProjectInternalPkgLogPath)); err != nil {
			return err
		}
	}

	return nil
}
