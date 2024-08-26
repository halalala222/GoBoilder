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

func newUserFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.DomainUserFileNae,
		template: domain.UserTemplate,
	}
}

func newErrorsFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.DomainErrorsFileName,
		template: domain.ErrorsTemplate,
	}
}

func getAllDomainFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		newUserFileBuilder(),
		newErrorsFileBuilder(),
	}
}

func (d *DomainBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range getAllDomainFileBuilder() {
		if err = fileBuild.build(filepath.Join(d.projectName, constants.ProjectDomainPkgPath)); err != nil {
			return err
		}
	}

	return nil
}
