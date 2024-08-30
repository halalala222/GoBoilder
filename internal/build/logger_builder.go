package build

import (
	"github.com/halalala222/GoBoilder/internal/template/logger"
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
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

func (l *LoggerBuilder) pkgLoggerTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(l.projectName, constants.ProjectLoggerPkgPath),
		Data:     nil,
	}
}

func (l *LoggerBuilder) internalLoggerTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(l.projectName, constants.ProjectInternalPkgLogPath),
		Data:     nil,
	}
}

func (l *LoggerBuilder) loggerLibraryTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(l.projectName, constants.ProjectLoggerPkgPath),
		Data:     nil,
	}
}

func (l *LoggerBuilder) newPkgLoggerFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  logger.GetPkgLoggerFileTemplateInfo(),
		buildInfo: l.pkgLoggerTemplateBuildInfo(),
	}
}

func (l *LoggerBuilder) newInternalLoggerFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  logger.GetInternalLogFileTemplateInfo(),
		buildInfo: l.internalLoggerTemplateBuildInfo(),
	}
}

func (l *LoggerBuilder) newLoggerLibraryFileBuilder() (*templateFileBuilder, error) {
	var (
		fileInfo *template.FileInfo
		err      error
	)

	if fileInfo, err = logger.GetLoggerLibraryFileTemplateInfo(l.library); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: l.loggerLibraryTemplateBuildInfo(),
	}, nil
}

func (l *LoggerBuilder) getAllPkgLoggerFileBuilder() ([]*templateFileBuilder, error) {
	var (
		fileBuilders = make([]*templateFileBuilder, 0, 2)
		err          error
		fileBuilder  *templateFileBuilder
	)

	if fileBuilder, err = l.newLoggerLibraryFileBuilder(); err != nil {
		return nil, err
	}

	fileBuilders = append(fileBuilders, l.newPkgLoggerFileBuilder())
	fileBuilders = append(fileBuilders, fileBuilder)

	return fileBuilders, nil
}

func (l *LoggerBuilder) getAllInternalPkgLoggerFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		l.newInternalLoggerFileBuilder(),
	}
}

func (l *LoggerBuilder) Build() error {
	var (
		err                     error
		allPkgLoggerFileBuilder []*templateFileBuilder
	)

	if allPkgLoggerFileBuilder, err = l.getAllPkgLoggerFileBuilder(); err != nil {
		return err
	}

	for _, fileBuilder := range allPkgLoggerFileBuilder {
		if err = fileBuilder.fileInfo.Build(fileBuilder.buildInfo); err != nil {
			return err
		}
	}

	for _, fileBuilder := range l.getAllInternalPkgLoggerFileBuilder() {
		if err = fileBuilder.fileInfo.Build(fileBuilder.buildInfo); err != nil {
			return err
		}
	}

	return nil
}
