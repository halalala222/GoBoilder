package cmd

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

func GetAppGoFileInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: appGoTmpl,
		FileName: constants.CmdAppFileName,
	}
}
