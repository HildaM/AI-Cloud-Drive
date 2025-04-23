package logger

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	gommlog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

// EchoAdapter adapts our zerolog logger to Echo's interface
type EchoAdapter struct {
	level  gommlog.Lvl // 日志级别
	prefix string      // 前缀
	header string      // 头部
}

// NewEchoLogger returns a new logger for Echo
func NewEchoLogger(cfg *Config) echo.Logger {
	// Map string level to gommon level
	var level gommlog.Lvl
	switch cfg.Level {
	case "debug":
		level = gommlog.DEBUG
	case "info":
		level = gommlog.INFO
	case "warn":
		level = gommlog.WARN
	case "error":
		level = gommlog.ERROR
	default:
		level = gommlog.INFO
	}

	return &EchoAdapter{
		level:  level,
		prefix: cfg.EchoPrefix,
		header: cfg.EchoHeader,
	}
}

// Output implements echo.Logger
func (a *EchoAdapter) Output() io.Writer {
	return logWriter
}

// SetOutput implements echo.Logger
func (a *EchoAdapter) SetOutput(w io.Writer) {
	SetOutput(w)
}

// Level implements echo.Logger
func (a *EchoAdapter) Level() gommlog.Lvl {
	return a.level
}

// SetLevel implements echo.Logger
func (a *EchoAdapter) SetLevel(lvl gommlog.Lvl) {
	a.level = lvl

	// Map gommon level to zerolog level
	var zerologLevel zerolog.Level
	switch lvl {
	case gommlog.DEBUG:
		zerologLevel = zerolog.DebugLevel
	case gommlog.INFO:
		zerologLevel = zerolog.InfoLevel
	case gommlog.WARN:
		zerologLevel = zerolog.WarnLevel
	case gommlog.ERROR:
		zerologLevel = zerolog.ErrorLevel
	default:
		zerologLevel = zerolog.InfoLevel
	}

	// Set zerolog global level
	zerolog.SetGlobalLevel(zerologLevel)
}

// SetHeader implements echo.Logger
func (a *EchoAdapter) SetHeader(h string) {
	a.header = h
}

// SetPrefix implements echo.Logger
func (a *EchoAdapter) SetPrefix(p string) {
	a.prefix = p
}

// Prefix implements echo.Logger
func (a *EchoAdapter) Prefix() string {
	return a.prefix
}

// Print implements echo.Logger
func (a *EchoAdapter) Print(args ...interface{}) {
	Info().Msg(fmt.Sprint(args...))
}

// Printf implements echo.Logger
func (a *EchoAdapter) Printf(format string, args ...interface{}) {
	Infof(format, args...)
}

// Printj implements echo.Logger
func (a *EchoAdapter) Printj(j gommlog.JSON) {
	event := Info()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Debug implements echo.Logger
func (a *EchoAdapter) Debug(args ...interface{}) {
	Debug().Msg(fmt.Sprint(args...))
}

// Debugf implements echo.Logger
func (a *EchoAdapter) Debugf(format string, args ...interface{}) {
	Debugf(format, args...)
}

// Debugj implements echo.Logger
func (a *EchoAdapter) Debugj(j gommlog.JSON) {
	event := Debug()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Info implements echo.Logger
func (a *EchoAdapter) Info(args ...interface{}) {
	Info().Msg(fmt.Sprint(args...))
}

// Infof implements echo.Logger
func (a *EchoAdapter) Infof(format string, args ...interface{}) {
	Infof(format, args...)
}

// Infoj implements echo.Logger
func (a *EchoAdapter) Infoj(j gommlog.JSON) {
	event := Info()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Warn implements echo.Logger
func (a *EchoAdapter) Warn(args ...interface{}) {
	Warn().Msg(fmt.Sprint(args...))
}

// Warnf implements echo.Logger
func (a *EchoAdapter) Warnf(format string, args ...interface{}) {
	Warnf(format, args...)
}

// Warnj implements echo.Logger
func (a *EchoAdapter) Warnj(j gommlog.JSON) {
	event := Warn()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Error implements echo.Logger
func (a *EchoAdapter) Error(args ...interface{}) {
	Error().Msg(fmt.Sprint(args...))
}

// Errorf implements echo.Logger
func (a *EchoAdapter) Errorf(format string, args ...interface{}) {
	Errorf(format, args...)
}

// Errorj implements echo.Logger
func (a *EchoAdapter) Errorj(j gommlog.JSON) {
	event := Error()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Fatal implements echo.Logger
func (a *EchoAdapter) Fatal(args ...interface{}) {
	Fatal().Msg(fmt.Sprint(args...))
}

// Fatalf implements echo.Logger
func (a *EchoAdapter) Fatalf(format string, args ...interface{}) {
	Fatalf(format, args...)
}

// Fatalj implements echo.Logger
func (a *EchoAdapter) Fatalj(j gommlog.JSON) {
	event := Fatal()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}

// Panic implements echo.Logger
func (a *EchoAdapter) Panic(args ...interface{}) {
	Panic().Msg(fmt.Sprint(args...))
}

// Panicf implements echo.Logger
func (a *EchoAdapter) Panicf(format string, args ...interface{}) {
	Panicf(format, args...)
}

// Panicj implements echo.Logger
func (a *EchoAdapter) Panicj(j gommlog.JSON) {
	event := Panic()
	for k, v := range j {
		event.Interface(k, v)
	}
	event.Send()
}
