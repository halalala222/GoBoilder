package logger

import _ "embed"

//go:embed pkg_logger.go.tmpl
var pkgLogger []byte

//go:embed internal_log.go.tmpl
var internalLog []byte

//go:embed zap_logger.go.tmpl
var zapLoggerTemplate []byte

//go:embed slog_logger.go.tmpl
var slogLoggerTemplate []byte
