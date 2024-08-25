package build

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/halalala222/GoBoilder/internal/constants"
)

type Builder interface {
	Build() error
}

func GenerateAllBuilder(options ...Option) []Builder {
	opts := getOptions(options)

	return []Builder{
		NewProjectBuilder(opts.projectName),
		NewLoggerBuilder(opts.projectName, opts.loggerLibrary, opts.modulePath),
	}
}

func getAllDir(projectName string) []string {
	return []string{
		filepath.Join(projectName, constants.ProjectLoggerPkgPath),
		filepath.Join(projectName, constants.ProjectInternalPkgLogPath),
	}
}

func AllDir(projectName string) error {
	var (
		err error
	)

	for _, dir := range getAllDir(projectName) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

type templateFileBuilder struct {
	fileName string
	template []byte
	data     any
}

func (t *templateFileBuilder) build(projectName string) error {
	var (
		err  error
		tmpl *template.Template
		file *os.File
	)

	if file, err = os.Create(filepath.Join(projectName, t.fileName)); err != nil {
		return err
	}

	if tmpl, err = template.New(t.fileName).Parse(string(t.template)); err != nil {
		return err
	}

	return tmpl.Execute(file, t.data)
}
