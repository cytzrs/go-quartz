package logger

import (
	"log"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Value

func init() {
	defaultLogger.Store(NewSimpleLogger(
		log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile), LevelInfo),
	)
}

// Default returns the default Logger.
func Default() Logger {
	return defaultLogger.Load().(Logger)
}

// SetDefault makes l the default Logger.
func SetDefault(l Logger) {
	defaultLogger.Store(l)
}

// Trace logs at LevelTrace.
func Trace(msg any) {
	Default().Trace(msg)
}

// Tracef logs at LevelTrace.
func Tracef(format string, args ...any) {
	Default().Tracef(format, args...)
}

// Debug logs at LevelDebug.
func Debug(msg any) {
	Default().Debug(msg)
}

// Debugf logs at LevelDebug.
func Debugf(format string, args ...any) {
	Default().Debugf(format, args...)
}

// Info logs at LevelInfo.
func Info(msg any) {
	Default().Info(msg)
}

// Infof logs at LevelInfo.
func Infof(format string, args ...any) {
	Default().Infof(format, args...)
}

// Warn logs at LevelWarn.
func Warn(msg any) {
	Default().Warn(msg)
}

// Warnf logs at LevelWarn.
func Warnf(format string, args ...any) {
	Default().Warnf(format, args...)
}

// Error logs at LevelError.
func Error(msg any) {
	Default().Error(msg)
}

// Errorf logs at LevelError.
func Errorf(format string, args ...any) {
	Default().Errorf(format, args...)
}

// Enabled reports whether the logger handles records at the given level.
func Enabled(level Level) bool {
	return Default().Enabled(level)
}
