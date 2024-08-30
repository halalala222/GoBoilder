package domain

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

func GetUserFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: userTemplate,
		FileName: constants.DomainUserFileName,
	}
}

func GetErrorsFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: errorsTemplate,
		FileName: constants.DomainErrorsFileName,
	}
}
