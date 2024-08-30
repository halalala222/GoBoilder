package logger

import _ "embed"

//go:embed pkg_logger.go.tmpl
var PkgLogger []byte

//go:embed internal_log.go.tmpl
var InternalLog []byte

//go:embed zap_logger.go.tmpl
var ZapLoggerTemplate []byte

//go:embed slog_logger.go.tmpl
var SlogLoggerTemplate []byte
