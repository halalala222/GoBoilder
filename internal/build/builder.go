package build

import (
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/halalala222/GoBoilder/internal/constants"
)

type Builder interface {
	Build() error
	String() string
}

func GenerateAllBuilder(options ...Option) []Builder {
	opts := getOptions(options)

	return []Builder{
		NewProjectBuilder(opts.projectName),
		NewLoggerBuilder(opts.projectName, opts.loggerLibrary, opts.modulePath),
		NewDomainBuilder(opts.projectName),
		NewConfigBuilder(opts.projectName, opts.modulePath, opts.db, opts.dbLibrary, opts.configFileType, opts.httpFramework),
		NewRepositoryBuilder(opts.projectName, opts.modulePath, opts.db, opts.dbLibrary),
		NewServiceBuilder(opts.projectName, opts.modulePath),
	}
}

func getAllDir(projectName string) []string {
	return []string{
		filepath.Join(projectName, constants.ProjectLoggerPkgPath),
		filepath.Join(projectName, constants.ProjectInternalPkgLogPath),
		filepath.Join(projectName, constants.ProjectDomainPkgPath),
		filepath.Join(projectName, constants.ProjectConfigPkgPath),
		filepath.Join(projectName, constants.ProjectRepositoryPkgPath),
		filepath.Join(projectName, constants.ProjectUserServicePkgPath),
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

func (t *templateFileBuilder) build(filePath string) error {
	var (
		err  error
		tmpl *template.Template
		file *os.File
	)

	if file, err = os.Create(filepath.Join(filePath, t.fileName)); err != nil {
		return err
	}

	if tmpl, err = template.New(t.fileName).Parse(string(t.template)); err != nil {
		return err
	}

	return tmpl.Execute(file, t.data)
}

func GoModInit(projectName string, modulePath string) error {
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = projectName

	return cmd.Run()
}

func GoModTidy(projectName string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName

	return cmd.Run()
}
