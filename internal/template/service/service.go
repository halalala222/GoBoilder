package service

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

func GetServiceFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: user,
		FileName: constants.ServiceFileName,
	}
}
