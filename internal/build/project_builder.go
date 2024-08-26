package build

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/project"
)

var _ Builder = &ProjectBuilder{}

type ProjectBuilder struct {
	projectName string
}

func (p *ProjectBuilder) String() string {
	return "ProjectBuilder"
}

func NewProjectBuilder(projectName string) *ProjectBuilder {
	return &ProjectBuilder{
		projectName: projectName,
	}
}

func newGitIgnoreFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.GitIgnoreFileName,
		template: project.GitIgnoreTemplate,
	}
}

func newREADMEFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.READEMEFileName,
		template: project.ReadmeTemplate,
	}
}

func newMakefileFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.MakefileFileName,
		template: project.MakefileTemplate,
	}
}

func getAllProjectFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		newGitIgnoreFileBuilder(),
		newREADMEFileBuilder(),
		newMakefileFileBuilder(),
	}
}

func (p *ProjectBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range getAllProjectFileBuilder() {
		if err = fileBuild.build(p.projectName); err != nil {
			return err
		}
	}

	return nil
}
