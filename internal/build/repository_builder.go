package build

import (
	"github.com/halalala222/GoBoilder/internal/template/repository"
	"os"
	"path/filepath"
	"strings"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var _ Builder = &RepositoryBuilder{}

type RepositoryBuilder struct {
	projectName string
	modulePath  string
	db          string
	dbLibrary   string
}

func (r *RepositoryBuilder) String() string {
	return "RepositoryBuilder"
}

func NewRepositoryBuilder(projectName, modulePath, db, dbLibrary string) *RepositoryBuilder {
	return &RepositoryBuilder{
		projectName: projectName,
		modulePath:  modulePath,
		db:          db,
		dbLibrary:   dbLibrary,
	}
}

func (r *RepositoryBuilder) repositoryTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: filepath.Join(r.projectName, constants.ProjectRepositoryPkgPath, strings.ToLower(r.db)),
		Data: &struct {
			ModulePath string
		}{
			ModulePath: r.modulePath,
		},
	}
}

func (r *RepositoryBuilder) newRepositoryFileBuilder() (*templateFileBuilder, error) {
	var (
		err      error
		fileInfo *template.FileInfo
	)

	if fileInfo, err = repository.GetRepositoryFileTemplateInfo(r.db, r.dbLibrary); err != nil {
		return nil, err
	}

	return &templateFileBuilder{
		fileInfo:  fileInfo,
		buildInfo: r.repositoryTemplateBuildInfo(),
	}, nil
}

func (r *RepositoryBuilder) Build() error {
	var (
		err                   error
		repositoryFileBuilder *templateFileBuilder
		dirPath               = filepath.Join(
			r.projectName,
			constants.ProjectRepositoryPkgPath,
			strings.ToLower(r.db),
		)
	)

	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	if repositoryFileBuilder, err = r.newRepositoryFileBuilder(); err != nil {
		return err
	}

	return repositoryFileBuilder.fileInfo.Build(repositoryFileBuilder.buildInfo)
}
