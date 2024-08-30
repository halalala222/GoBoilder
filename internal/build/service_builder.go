package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/service"
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

func (s *ServiceBuilder) newServiceFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.ServiceFileName,
		template: service.User,
		data: &struct {
			ModulePath string
		}{
			ModulePath: s.modulePath,
		},
	}
}

func (s *ServiceBuilder) Build() error {
	return s.newServiceFileBuilder().build(filepath.Join(s.projectName, constants.ProjectUserServicePkgPath))
}
