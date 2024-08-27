package constants

import "errors"

var (
	ErrLoggerLibraryNotSupported = errors.New("logger library not supported")
	ErrHTTPFrameNotSupported     = errors.New("http frame not supported")
	ErrQuit                      = errors.New("quit")

	ErrProjectNameEmpty             = errors.New("project name cannot be empty")
	ErrProjectNameExists            = errors.New("project name already exists")
	ErrProjectNameStartOrEndWithDot = errors.New("project name cannot start or end with dot")
	ErrInvalidASCIICharacters       = errors.New("invalid ASCII characters")

	ErrModulePathEmpty = errors.New("module path cannot be empty")
)
