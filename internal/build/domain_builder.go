package build

import (
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/domain"
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
		fileName: constants.DomainUserFileNae,
		template: domain.UserTemplate,
	}
}

func (d *DomainBuilder) newErrorsFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.DomainErrorsFileName,
		template: domain.ErrorsTemplate,
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
		if err = fileBuild.build(filepath.Join(d.projectName, constants.ProjectDomainPkgPath)); err != nil {
			return err
		}
	}

	return nil
}
