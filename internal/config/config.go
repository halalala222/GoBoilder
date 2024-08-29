package config

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/config"
)

type FileTemplate interface {
	Build() []byte
}

var supportedFileTemplates = map[string]FileTemplate{
	constants.YAMLConfigFileType: &YAMLTemplate{},
	constants.TOMLConfigFileType: &TOMLTemplate{},
	constants.JSONConfigFileType: &JSONTemplate{},
	constants.ENVConfigFileType:  &ENVTemplate{},
}

func GetConfigTemplate() []byte {
	return config.Template
}

// GetConfigFileTemplate returns the FileTemplate for the given fileType
func GetConfigFileTemplate(fileType string) (FileTemplate, error) {
	var (
		template FileTemplate
		ok       bool
	)

	if template, ok = supportedFileTemplates[fileType]; !ok {
		return nil, constants.ErrUnsupportedFileType
	}

	return template, nil
}

// GetConfigFileName returns the name of the config file for the given fileType
func GetConfigFileName(fileType string) (string, error) {
	configFileType2ConfigFileName := map[string]string{
		constants.YAMLConfigFileType: constants.YAMLConfigFileName,
		constants.TOMLConfigFileType: constants.TOMLConfigFileName,
		constants.JSONConfigFileType: constants.JSONConfigFileName,
		constants.ENVConfigFileType:  constants.ENVConfigFileName,
	}

	if fileName, ok := configFileType2ConfigFileName[fileType]; ok {
		return fileName, nil
	}

	return "", constants.ErrUnsupportedFileType
}

// GetSupportedConfigFileTypes returns the supported file types
func GetSupportedConfigFileTypes() []string {
	var fileTypes = make([]string, 0, len(supportedFileTemplates))

	for fileType := range supportedFileTemplates {
		fileTypes = append(fileTypes, fileType)
	}

	return fileTypes
}
