package template

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/logger"
)

var supportedLoggerLibraries = map[string]*FileInfo{
	constants.ZapLoggerLibrary: {
		Template: logger.SlogLoggerTemplate,
		FileName: constants.ZapLoggerFileName,
	},
	constants.SlogLoggerLibrary: {
		Template: logger.ZapLoggerTemplate,
		FileName: constants.SlogLoggerFileName,
	},
}

// GetInternalLogFileTemplateInfo returns the internal log file template.
func GetInternalLogFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: logger.InternalLog,
		FileName: constants.InternalLogFileName,
	}
}

// GetPkgLoggerFileTemplateInfo returns the package logger file template.
func GetPkgLoggerFileTemplateInfo() *FileInfo {
	return &FileInfo{
		Template: logger.PkgLogger,
		FileName: constants.PkgLoggerFileName,
	}
}

// GetLoggerLibraryFileTemplateInfo returns the LoggerTemplateInfo for the given library.
func GetLoggerLibraryFileTemplateInfo(library string) (*FileInfo, error) {
	var (
		loggerTemplateInfo *FileInfo
		ok                 bool
	)

	if loggerTemplateInfo, ok = supportedLoggerLibraries[library]; !ok {
		return nil, constants.ErrLoggerLibraryNotSupported
	}

	return loggerTemplateInfo, nil
}
