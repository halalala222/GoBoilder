package logger

import (
	_ "embed"

	"github.com/halalala222/GoBoilder/internal/template/logger"
)

var _ LibraryTemplate = &SlogLoggerTemplate{}

// SlogLoggerTemplate used to build the slog logger library.
type SlogLoggerTemplate struct{}

func (s *SlogLoggerTemplate) Build() []byte {
	return logger.SlogLoggerTemplate
}
