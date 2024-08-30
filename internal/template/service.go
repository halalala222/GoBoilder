package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/service"
)

func GetServiceFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: service.User,
		FileName: constants.ServiceFileName,
	}
}
