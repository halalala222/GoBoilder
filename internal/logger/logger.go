package logger

import (
	_ "embed"

	"github.com/halalala222/GoBoilder/internal/constants"
	"github.com/halalala222/GoBoilder/internal/template/logger"
)

// TemplateLogger returns a byte slice that represents
// the logger interface template.
func TemplateLogger() []byte {
	return logger.Logger
}

func TemplateLog() []byte {
	return logger.Log
}

// LibraryTemplate is an interface that defines the Build method.
// to build the logger library.
type LibraryTemplate interface {
	Build() []byte
}

var supportedLibraries = map[string]LibraryTemplate{
	constants.ZapLoggerLibrary:  &ZapLoggerTemplate{},
	constants.SlogLoggerLibrary: &SlogLoggerTemplate{},
}

// GetLibraryTemplate returns the LibraryTemplate for the given library.
func GetLibraryTemplate(library string) (LibraryTemplate, error) {
	var (
		libraryTemplate LibraryTemplate
		ok              bool
	)

	if libraryTemplate, ok = supportedLibraries[library]; !ok {
		return nil, constants.ErrLoggerLibraryNotSupported
	}

	return libraryTemplate, nil
}

func GetAllSupportedLibraries() []string {
	var libraries = make([]string, 0, len(supportedLibraries))

	for library := range supportedLibraries {
		libraries = append(libraries, library)
	}

	return libraries
}
