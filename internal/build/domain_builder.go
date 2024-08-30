package build

import (
	"github.com/halalala222/GoBoilder/internal/template/domain"
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var _ Builder = &DomainBuilder{}

type DomainBuilder struct {
	projectName string
}

func (d *DomainBuilder) String() string {
	return "DomainBuilder"
}

func NewDomainBuilder(projectName string) *DomainBuilder {
	return &DomainBuilder{
		projectName: projectName,
	}
}

func (d *DomainBuilder) newUserFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo: domain.GetUserFileTemplateInfo(),
		buildInfo: &template.BuildInfo{
			FilePath: filepath.Join(d.projectName, constants.ProjectDomainPkgPath),
			Data:     nil,
		},
	}
}

func (d *DomainBuilder) newErrorsFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo: domain.GetErrorsFileTemplateInfo(),
		buildInfo: &template.BuildInfo{
			FilePath: filepath.Join(d.projectName, constants.ProjectDomainPkgPath),
			Data:     nil,
		},
	}
}

func (d *DomainBuilder) getAllDomainFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		d.newUserFileBuilder(),
		d.newErrorsFileBuilder(),
	}
}

func (d *DomainBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range d.getAllDomainFileBuilder() {
		if err = fileBuild.fileInfo.Build(fileBuild.buildInfo); err != nil {
			return err
		}
	}

	return nil
}
