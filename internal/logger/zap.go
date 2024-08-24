package logger

import (
	_ "embed"

	"github.com/halalala222/GoBoilder/internal/template/logger"
)

var _ LibraryTemplate = &ZapLoggerTemplate{}

// ZapLoggerTemplate used to build the zap logger library.
type ZapLoggerTemplate struct{}

func (z *ZapLoggerTemplate) Build() []byte {
	return logger.ZapLoggerTemplate
}
