package logger

import (
	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template"
)

var supportedLoggerLibraries = map[string]*template.FileInfo{
	constants.ZapLoggerLibrary: {
		Template: slogLoggerTemplate,
		FileName: constants.ZapLoggerFileName,
	},
	constants.SlogLoggerLibrary: {
		Template: zapLoggerTemplate,
		FileName: constants.SlogLoggerFileName,
	},
}

// GetInternalLogFileTemplateInfo returns the internal log file template.
func GetInternalLogFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: internalLog,
		FileName: constants.InternalLogFileName,
	}
}

// GetPkgLoggerFileTemplateInfo returns the package logger file template.
func GetPkgLoggerFileTemplateInfo() *template.FileInfo {
	return &template.FileInfo{
		Template: pkgLogger,
		FileName: constants.PkgLoggerFileName,
	}
}

// GetLoggerLibraryFileTemplateInfo returns the LoggerTemplateInfo for the given library.
func GetLoggerLibraryFileTemplateInfo(library string) (*template.FileInfo, error) {
	var (
		loggerTemplateInfo *template.FileInfo
		ok                 bool
	)

	if loggerTemplateInfo, ok = supportedLoggerLibraries[library]; !ok {
		return nil, constants.ErrLoggerLibraryNotSupported
	}

	return loggerTemplateInfo, nil
}
