## For logger
> contains the logger implementation for the project.

1. generate pkg/logger/logger.go define the logger interface for the project. While implementing the logger interface, you can use any logger library like zap, slog, etc.
```go
package log

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Sync() error
}

type Field struct {
	Key   string
	Value any
}
```

2. generate pkg/logger/zap.go or pkg/logger/slog.go or any other logger implementation file. 

zap_logger.go
```go
package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = &ZapLogger{}

type ZapLogger struct {
	logger *zap.Logger
}

var (
	highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
)

func NewZapLogger(logPath *Path, options ...zap.Option) *ZapLogger {
	zapCores := make([]zapcore.Core, 0, 4)
	zapCores = append(zapCores, getConsoleStdOutZapCore())
	zapCores = append(zapCores, getConsoleStdErrZapCore())
	if !logPath.checkInfoPathIsEmpty() {
		zapCores = append(zapCores, getFileInfoCore(logPath.InfoPath))
	}
	if !logPath.checkErrorPathIsEmpty() {
		zapCores = append(zapCores, getFileErrorCore(logPath.ErrorPath))
	}

	logger := zap.New(
		zapcore.NewTee(
			zapCores...,
		),
		options...,
	)

	return &ZapLogger{
		logger: logger,
	}
}

func getConsoleStdOutZapCore() zapcore.Core {
	return zapcore.NewCore(
		getConsoleEncoder(),
		zapcore.Lock(os.Stdout),
		lowPriority,
	)
}

func getConsoleStdErrZapCore() zapcore.Core {
	return zapcore.NewCore(
		getConsoleEncoder(),
		zapcore.Lock(os.Stderr),
		highPriority,
	)
}

func getFileInfoCore(infoPath string) zapcore.Core {
	return zapcore.NewCore(
		getJSONEncoder(),
		getLogWriteSyncer(infoPath),
		lowPriority,
	)
}

func getFileErrorCore(errorPath string) zapcore.Core {
	return zapcore.NewCore(
		getJSONEncoder(),
		getLogWriteSyncer(errorPath),
		highPriority,
	)
}

func getLogWriteSyncer(logFilePath string) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *ZapLogger) Logger() *zap.Logger {
	return l.logger
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, l.toZapFields(fields)...)
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

func (l *ZapLogger) toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
```

slog_logger.go
```go
package log

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"golang.org/x/exp/slog"
)

var _ Logger = &SLogLogger{}

type SLogLogger struct {
	logger *slog.Logger
}

const (
	levelTrace = slog.Level(-8)
	levelFatal = slog.Level(12)
	levelPanic = slog.Level(16)
)

var (
	levelNames = map[slog.Leveler]string{
		levelTrace: "TRACE",
		levelFatal: "FATAL",
		levelPanic: "PANIC",
	}
	replaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			level := a.Value.Any().(slog.Level)
			levelLabel, exists := levelNames[level]
			if !exists {
				levelLabel = level.String()
			}

			a.Value = slog.StringValue(levelLabel)
		}

		return a
	}
)

func NewSlogLogger(path string) *SLogLogger {
	var handler = getDevHandler(path)
	if os.Getenv("APP_ENV") == "prod" {
		handler = getProdHandler(path)
	}
	slogLogger := slog.New(handler)
	return &SLogLogger{
		logger: slogLogger,
	}
}

func getFileWriter(path string) io.Writer {
	if len(path) == 0 {
		return os.Stdout
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return io.MultiWriter(os.Stdout, lumberJackLogger)
}

func getProdOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: replaceAttr,
	}
}

func getDevOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: replaceAttr,
	}
}

func getProdHandler(path string) slog.Handler {
	return slog.NewJSONHandler(getFileWriter(path), getProdOpts())
}

func getDevHandler(path string) slog.Handler {
	return slog.NewTextHandler(getFileWriter(path), getDevOpts())
}

func (s *SLogLogger) Debug(msg string, fields ...Field) {
	s.logger.Debug(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Info(msg string, fields ...Field) {
	s.logger.Info(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Warn(msg string, fields ...Field) {
	s.logger.Warn(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Error(msg string, fields ...Field) {
	s.logger.Error(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Fatal(msg string, fields ...Field) {
	s.logger.Log(nil, levelFatal, msg, s.toSlogAttrs(fields)...)
	os.Exit(1)
}

func (s *SLogLogger) Panic(msg string, fields ...Field) {
	s.logger.Log(nil, levelPanic, msg, s.toSlogAttrs(fields)...)
	panic(msg)
}

func (s *SLogLogger) Sync() error {
	return nil
}

func (s *SLogLogger) toSlogAttrs(fields []Field) []any {
	slogAttrs := make([]any, 0, len(fields))
	for _, f := range fields {
		slogAttrs = append(slogAttrs, slog.Any(f.Key, f.Value))
	}
	return slogAttrs
}
```

3. generate internal/pkg/log/log.go file to initialize the logger.Provider and register the logger implementation.Project can use this file to get the logger instance.
```go
package log

import (
	logger "{{.ModulePath}}/pkg/log"
)

var log logger.Logger

func RegisterLog(logger logger.Logger) {
	log = logger
}

func Log() logger.Logger {
	if log == nil {
		panic("implement not found for interface Logger, please register")
	}
	return log
}
```