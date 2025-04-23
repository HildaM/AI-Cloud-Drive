package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

// Global logger instance
var zlog zerolog.Logger

// Global writer
var logWriter io.Writer

// 将字符串级别转换为 zerolog 级别
func parseLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// Init 初始化日志系统
func Init(cfg Config) error {
	// 配置校验
	cfg.Check()

	// 设置全局日志级别
	zerolog.SetGlobalLevel(parseLevel(cfg.Level))

	// 创建 writers
	var writers []io.Writer

	// 配置控制台输出
	if cfg.ConsoleOutput {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		writers = append(writers, consoleWriter)
	}

	// 配置文件输出
	if cfg.FileOutput {
		// 创建目录（如果不存在）
		dir := filepath.Dir(cfg.FilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建日志目录失败: %w", err)
		}

		// 打开日志文件
		file, err := os.OpenFile(cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("打开日志文件失败: %w", err)
		}

		writers = append(writers, file)
	}

	// 确保至少有一个输出
	if len(writers) == 0 {
		return fmt.Errorf("未配置任何日志输出")
	}

	// 创建多writer（如有必要）
	if len(writers) == 1 {
		logWriter = writers[0]
	} else {
		logWriter = zerolog.MultiLevelWriter(writers...)
	}

	// 创建日志实例
	logContext := zerolog.New(logWriter).With().Timestamp()
	if cfg.WithCaller {
		logContext = logContext.Caller()
	}
	zlog = logContext.Logger()
	return nil
}

// Debug logs a debug message
func Debug() *zerolog.Event {
	return zlog.Debug()
}

// Info logs an info message
func Info() *zerolog.Event {
	return zlog.Info()
}

// Warn logs a warning message
func Warn() *zerolog.Event {
	return zlog.Warn()
}

// Error logs an error message
func Error() *zerolog.Event {
	return zlog.Error()
}

// Fatal logs a fatal message
func Fatal() *zerolog.Event {
	return zlog.Fatal()
}

// Panic logs a panic message
func Panic() *zerolog.Event {
	return zlog.Panic()
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	zlog.Debug().Msgf(format, args...)
}

// Infof logs a formatted info message
func Infof(format string, args ...interface{}) {
	zlog.Info().Msgf(format, args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	zlog.Warn().Msgf(format, args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	zlog.Error().Msgf(format, args...)
}

// Fatalf logs a formatted fatal message
func Fatalf(format string, args ...interface{}) {
	zlog.Fatal().Msgf(format, args...)
}

// Panicf logs a formatted panic message
func Panicf(format string, args ...interface{}) {
	zlog.Panic().Msgf(format, args...)
}

// SetOutput sets the log output
func SetOutput(w io.Writer) {
	logWriter = w
	zlog = zlog.Output(w)
}

// WithField adds a field to the log
func WithField(key string, value interface{}) *zerolog.Event {
	return Info().Interface(key, value)
}

// WithFields adds multiple fields to the log
func WithFields(fields map[string]interface{}) *zerolog.Event {
	event := zlog.Info()
	for k, v := range fields {
		event.Interface(k, v)
	}
	return event
}
