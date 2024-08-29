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

func (p *ProjectBuilder) newGitIgnoreFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.GitIgnoreFileName,
		template: project.GitIgnoreTemplate,
	}
}

func (p *ProjectBuilder) newREADMEFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.READEMEFileName,
		template: project.ReadmeTemplate,
	}
}

func (p *ProjectBuilder) newMakefileFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileName: constants.MakefileFileName,
		template: project.MakefileTemplate,
	}
}

func (p *ProjectBuilder) getAllProjectFileBuilder() []*templateFileBuilder {
	return []*templateFileBuilder{
		p.newGitIgnoreFileBuilder(),
		p.newREADMEFileBuilder(),
		p.newMakefileFileBuilder(),
	}
}

func (p *ProjectBuilder) Build() error {
	var (
		err error
	)

	for _, fileBuild := range p.getAllProjectFileBuilder() {
		if err = fileBuild.build(p.projectName); err != nil {
			return err
		}
	}

	return nil
}
