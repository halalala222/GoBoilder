package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var _ Builder = &ServiceBuilder{}

type ServiceBuilder struct {
	projectName string
	modulePath  string
}

func (s *ServiceBuilder) String() string {
	return "ServiceBuilder"
}

func NewServiceBuilder(projectName, modulePath string) *ServiceBuilder {
	return &ServiceBuilder{
		projectName: projectName,
		modulePath:  modulePath,
	}
}

func (s *ServiceBuilder) serviceFileBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(s.projectName, constants.ProjectUserServicePkgPath),
		Data:     nil,
	}
}

func (s *ServiceBuilder) newServiceFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  template.GetServiceFileTemplateInfo(),
		buildInfo: s.serviceFileBuildInfo(),
	}
}

func (s *ServiceBuilder) Build() error {
	return s.newServiceFileBuilder().fileInfo.Build(s.newServiceFileBuilder().buildInfo)
}
