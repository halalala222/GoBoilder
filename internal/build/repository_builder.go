package build

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
	"github.com/halalala222/GoBoilder/internal/template/repository"
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

func (r *RepositoryBuilder) getRepositoryDir() string {
	return filepath.Join(r.projectName, constants.ProjectRepositoryPkgPath, strings.ToLower(r.db))
}

func (r *RepositoryBuilder) repositoryTemplateBuildInfo() *template.BuildInfo {
	return &template.BuildInfo{
		FilePath: r.getRepositoryDir(),
		Data: &struct {
			ModulePath string
			DB         string
		}{
			ModulePath: r.modulePath,
			DB:         strings.ToLower(r.db),
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
		dirPath               = r.getRepositoryDir()
	)

	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	if repositoryFileBuilder, err = r.newRepositoryFileBuilder(); err != nil {
		return err
	}

	return repositoryFileBuilder.fileInfo.Build(repositoryFileBuilder.buildInfo)
}
