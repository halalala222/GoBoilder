package build

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/halalala222/GoBoilder/internal/constants"
	templatePkg "github.com/halalala222/GoBoilder/internal/template"
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
		NewConfigBuilder(opts.projectName, opts.modulePath, opts.db, opts.dbLibrary, opts.configFileType, opts.httpFramework, opts.loggerLibrary),
		NewRepositoryBuilder(opts.projectName, opts.modulePath, opts.db, opts.dbLibrary),
		NewServiceBuilder(opts.projectName, opts.modulePath),
		NewRestBuilder(opts.projectName, opts.modulePath, opts.httpFramework),
		NewAppBuilder(opts.projectName, opts.modulePath, opts.httpFramework, opts.db, opts.dbLibrary),
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
		filepath.Join(projectName, constants.ProjectRestPkgPath),
		filepath.Join(projectName, constants.ProjectCmdPkgPath),
		filepath.Join(projectName, constants.ProjectLogsPath),
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
	fileInfo  *templatePkg.FileInfo
	buildInfo *templatePkg.BuildInfo
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
