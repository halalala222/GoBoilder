package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
	"github.com/halalala222/GoBoilder/internal/template/rest"
)

var _ Builder = &RestBuilder{}

type RestBuilder struct {
	projectName   string
	modulePath    string
	httpFramework string
}

func (r *RestBuilder) String() string {
	return "RestBuilder"
}

func NewRestBuilder(projectName, modulePath, httpFramework string) *RestBuilder {
	return &RestBuilder{
		projectName:   projectName,
		modulePath:    modulePath,
		httpFramework: httpFramework,
	}
}

func (r *RestBuilder) restHandlerFileBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(r.projectName, constants.ProjectRestPkgPath),
		Data: &struct {
			ModulePath    string
			HTTPFramework string
		}{
			ModulePath:    r.modulePath,
			HTTPFramework: r.httpFramework,
		},
	}
}

func (r *RestBuilder) getRestHandlerFileBuilder() (*templateFileBuilder, error) {
	var (
		err      error
		fileInfo *template.FileInfo
	)

	if fileInfo, err = rest.GetRestHandlerFileTemplateInfo(r.httpFramework); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: r.restHandlerFileBuildInfo(),
	}, nil
}

func (r *RestBuilder) Build() error {
	var (
		err                    error
		restHandlerFileBuilder *templateFileBuilder
	)

	if restHandlerFileBuilder, err = r.getRestHandlerFileBuilder(); err != nil {
		return err
	}

	return restHandlerFileBuilder.fileInfo.Build(restHandlerFileBuilder.buildInfo)
}
