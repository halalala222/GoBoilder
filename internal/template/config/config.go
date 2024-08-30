package config

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var supportedConfigFileType = map[string]*template.FileInfo{
	constants.YAMLConfigFileType: {
		Template: yamlTemplate,
		FileName: constants.YAMLConfigFileName,
	},
	constants.TOMLConfigFileType: {
		Template: tomlTemplate,
		FileName: constants.TOMLConfigFileName,
	},
	constants.JSONConfigFileType: {
		Template: jsonTemplate,
		FileName: constants.JSONConfigFileName,
	},
	constants.ENVConfigFileType: {
		Template: envTemplate,
		FileName: constants.ENVConfigFileName,
	},
}

// GetConfigFileTemplateInfo returns the FileInfo for the config file template.
func GetConfigFileTemplateInfo(configFileType string) (*template.FileInfo, error) {

	if fileInfo, ok := supportedConfigFileType[configFileType]; ok {
		return fileInfo, nil
	}

	return nil, constants.ErrUnsupportedFileType
}

// GetSupportedConfigFileTypes returns the supported file types
func GetSupportedConfigFileTypes() []string {
	var fileTypes = make([]string, 0, len(supportedConfigFileType))

	for fileType := range supportedConfigFileType {
		fileTypes = append(fileTypes, fileType)
	}

	return fileTypes
}

// GetConfigLoaderFileTemplateInfo returns the FileInfo for the config loader file template.
func GetConfigLoaderFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: loaderTemplate,
		FileName: constants.ConfigLoaderFileName,
	}
}
