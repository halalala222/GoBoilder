package logger

import _ "embed"

//go:embed logger.go.tmpl
var Logger []byte

//go:embed log.go.tmpl
var Log []byte

//go:embed zap_logger.go.tmpl
var ZapLoggerTemplate []byte

//go:embed slog_logger.go.tmpl
var SlogLoggerTemplate []byte
