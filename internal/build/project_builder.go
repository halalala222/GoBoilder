package build

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/project"
)

var _ Builder = &ProjectBuilder{}

type ProjectBuilder struct {
	projectName string
}

func NewProjectBuilder(projectName string) *ProjectBuilder {
	return &ProjectBuilder{
		projectName: projectName,
	}
}

type projectFileBuild struct {
	fileName string
	template []byte
}

func newGitIgnoreFileBuild() *projectFileBuild {
	return &projectFileBuild{
		fileName: constants.GitIgnoreFileName,
		template: project.GitIgnoreTemplate,
	}
}

func newREADMEFileBuild() *projectFileBuild {
	return &projectFileBuild{
		fileName: constants.READEMEFileName,
		template: project.ReadmeTemplate,
	}
}

func newMakefileFileBuild() *projectFileBuild {
	return &projectFileBuild{
		fileName: constants.MakefileFileName,
		template: project.MakefileTemplate,
	}
}

func getAllProjectFileBuild() []*projectFileBuild {
	return []*projectFileBuild{
		newGitIgnoreFileBuild(),
		newREADMEFileBuild(),
		newMakefileFileBuild(),
	}
}

func (p *projectFileBuild) build(projectName string) error {
	var (
		err  error
		tmpl *template.Template
		file *os.File
	)

	if file, err = os.Create(filepath.Join(projectName, p.fileName)); err != nil {
		return err
	}

	if tmpl, err = template.New(p.fileName).Parse(string(p.template)); err != nil {
		return err
	}

	return tmpl.Execute(file, nil)
}

func (p *ProjectBuilder) Build() error {
	var (
		err error
	)

	if err = os.Mkdir(p.projectName, os.ModePerm); err != nil {
		return err
	}

	for _, fileBuild := range getAllProjectFileBuild() {
		if err = fileBuild.build(p.projectName); err != nil {
			return err
		}
	}

	return nil
}
