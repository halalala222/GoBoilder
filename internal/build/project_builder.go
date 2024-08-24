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

func (p *ProjectBuilder) buildGitIgnore() error {
	var (
		err  error
		tmpl *template.Template
		file *os.File
	)

	if file, err = os.Create(filepath.Join(p.projectName, constants.GitIgnoreFileName)); err != nil {
		return err
	}

	if tmpl, err = template.New(".gitignore").Parse(string(project.GitIgnoreTemplate)); err != nil {
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

	if err = p.buildGitIgnore(); err != nil {
		return err
	}

	return nil
}
