package rest

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

func getHTTPFramework2FileTemplateInfo() map[string]*template.FileInfo {
	return map[string]*template.FileInfo{
		constants.HTTPFrameworkGin: {
			Template: restGinTemplate,
			FileName: constants.RestUserFileName,
		},
		constants.HTTPFrameworkEcho: {
			Template: restEchoTemplate,
			FileName: constants.RestUserFileName,
		},
		constants.HTTPFrameworkChi: {
			Template: restChiTemplate,
			FileName: constants.RestUserFileName,
		},
		constants.HTTPFrameworkFiber: {
			Template: restFiberTemplate,
			FileName: constants.RestUserFileName,
		},
	}
}

// GetRestHandlerFileTemplateInfo returns the FileInfo for the rest handler file template.
func GetRestHandlerFileTemplateInfo(httpFramework string) (*template.FileInfo, error) {
	var (
		httpFramework2FileTemplateInfo = getHTTPFramework2FileTemplateInfo()
	)

	if fileInfo, ok := httpFramework2FileTemplateInfo[httpFramework]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrUnsupportedHTTPFramework
}
