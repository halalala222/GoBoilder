package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/domain"
)

func GetUserFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: domain.UserTemplate,
		FileName: constants.DomainUserFileName,
	}
}

func GetErrorsFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: domain.ErrorsTemplate,
		FileName: constants.DomainErrorsFileName,
	}
}
