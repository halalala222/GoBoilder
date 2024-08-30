package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/project"
)

func GetGitIgnoreFileInfo() *FileInfo {
	return &FileInfo{
		Template: project.GitIgnoreTemplate,
		FileName: constants.GitIgnoreFileName,
	}
}

func GetREADMEFileInfo() *FileInfo {
	return &FileInfo{
		Template: project.ReadmeTemplate,
		FileName: constants.READEMEFileName,
	}
}

func GetMakefileFileInfo() *FileInfo {
	return &FileInfo{
		Template: project.MakefileTemplate,
		FileName: constants.MakefileFileName,
	}
}
