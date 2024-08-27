package constants

import "errors"

var (
	ErrLoggerLibraryNotSupported = errors.New("logger library not supported")
	ErrHTTPFrameNotSupported     = errors.New("http frame not supported")
	ErrQuit                      = errors.New("quit")

	ErrProjectNameEmpty                                 = errors.New("project name cannot be empty")
	ErrProjectNameExists                                = errors.New("project name already exists")
	ErrProjectNameStartOrEndWithDot                     = errors.New("project name cannot start or end with dot")
	ErrProjectModulePathPrefixStartOrEndWithSlash       = errors.New("module path prefix cannot start or end with slash")
	ErrProjectModulePathPrefixContainsDoubleSlash       = errors.New("module path prefix cannot contain double slash")
	ErrProjectModulePathContainsWindowsReservedFileName = errors.New("module path contains windows reserved file name")
	ErrProjectModulePathEndWithTildeFollowedByDigits    = errors.New("module path end with tilde followed by digits")
	ErrInvalidASCIICharacters                           = errors.New("invalid ASCII characters")

	ErrModulePathEmpty = errors.New("module path cannot be empty")
)
