package build

import (
	"github.com/halalala222/GoBoilder/internal/template"
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

func (p *ProjectBuilder) gitIgnoreTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: p.projectName,
		Data:     nil,
	}
}

func (p *ProjectBuilder) readmeTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: p.projectName,
		Data:     nil,
	}
}

func (p *ProjectBuilder) makefileTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: p.projectName,
		Data:     nil,
	}
}

func (p *ProjectBuilder) newGitIgnoreFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  project.GetGitIgnoreFileInfo(),
		buildInfo: p.gitIgnoreTemplateBuildInfo(),
	}
}

func (p *ProjectBuilder) newREADMEFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  project.GetREADMEFileInfo(),
		buildInfo: p.readmeTemplateBuildInfo(),
	}
}

func (p *ProjectBuilder) newMakefileFileBuilder() *templateFileBuilder {
	return &templateFileBuilder{
		fileInfo:  project.GetMakefileFileInfo(),
		buildInfo: p.makefileTemplateBuildInfo(),
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
		if err = fileBuild.fileInfo.Build(fileBuild.buildInfo); err != nil {
			return err
		}
	}

	return nil
}
