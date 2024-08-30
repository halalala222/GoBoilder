package project

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

func GetGitIgnoreFileInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: gitIgnoreTemplate,
		FileName: constants.GitIgnoreFileName,
	}
}

func GetREADMEFileInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: readmeTemplate,
		FileName: constants.READEMEFileName,
	}
}

func GetMakefileFileInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: makefileTemplate,
		FileName: constants.MakefileFileName,
	}
}
