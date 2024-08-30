package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config"
)

var supportedConfigFileType = map[string]*FileInfo{
	constants.YAMLConfigFileType: {
		Template: config.YAMLTemplate,
		FileName: constants.YAMLConfigFileName,
	},
	constants.TOMLConfigFileType: {
		Template: config.TOMLTemplate,
		FileName: constants.TOMLConfigFileName,
	},
	constants.JSONConfigFileType: {
		Template: config.JSONTemplate,
		FileName: constants.JSONConfigFileName,
	},
	constants.ENVConfigFileType: {
		Template: config.ENVTemplate,
		FileName: constants.ENVConfigFileName,
	},
}

// GetConfigFileTemplateInfo returns the FileInfo for the config file template.
func GetConfigFileTemplateInfo(configFileType string) (*FileInfo, error) {

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
func GetConfigLoaderFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: config.Template,
		FileName: constants.ConfigLoaderFileName,
	}
}
