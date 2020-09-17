package log

import (
	"go.uber.org/zap"
)

func SetDebug() {

	defaultLogger.SetDebug()

}

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	With(fields ...interface{}) Logger
}

// Debuga uses fmt.Sprint to construct and log a message at debug level.
// This call is a wrapper around [Sugaredlogger.Debug](https://godoc.org/go.uber.org/zap#Sugaredlogger.Debug)
func Debug(args ...interface{}) {
	defaultLogger.sLog.Debug(args...)
}

// Debugf uses fmt.Sprintf to construct and log a message at debug level.
// This call is a wrapper around [Sugaredlogger.Debugf](https://godoc.org/go.uber.org/zap#Sugaredlogger.Debugf)
func Debugf(template string, args ...interface{}) {
	defaultLogger.sLog.Debugf(template, args...)
}

// Debugw logs a message at debug level with some additional context.
// This call is a wrapper around [Sugaredlogger.Debugw](https://godoc.org/go.uber.org/zap#Sugaredlogger.Debugw)
func Debugw(msg string, keysAndValues ...interface{}) {
	defaultLogger.sLog.Debugw(msg, keysAndValues...)
}

// DebugEnabled returns whether output of messages at the debug level is currently enabled.
func DebugEnabled() bool {
	return defaultLogger.Core().Enabled(zap.DebugLevel)
}

// Errora uses fmt.Sprint to construct and log a message at error level.
// This call is a wrapper around [Sugaredlogger.Error](https://godoc.org/go.uber.org/zap#Sugaredlogger.Error)
func Error(args ...interface{}) {
	defaultLogger.sLog.Error(args...)
}

// Errorf uses fmt.Sprintf to construct and log a message at error level.
// This call is a wrapper around [Sugaredlogger.Errorf](https://godoc.org/go.uber.org/zap#Sugaredlogger.Errorf)
func Errorf(template string, args ...interface{}) {
	defaultLogger.sLog.Errorf(template, args...)
}

// Errorw logs a message at error level with some additional context.
// This call is a wrapper around [Sugaredlogger.Errorw](https://godoc.org/go.uber.org/zap#Sugaredlogger.Errorw)
func Errorw(msg string, keysAndValues ...interface{}) {
	defaultLogger.sLog.Errorw(msg, keysAndValues...)
}

// ErrorEnabled returns whether output of messages at the error level is currently enabled.
func ErrorEnabled() bool {
	return defaultLogger.Core().Enabled(zap.ErrorLevel)
}

// Warna uses fmt.Sprint to construct and log a message at warn level.
// This call is a wrapper around [Sugaredlogger.Warn](https://godoc.org/go.uber.org/zap#Sugaredlogger.Warn)
func Warn(args ...interface{}) {
	defaultLogger.sLog.Warn(args...)
}

// Warnf uses fmt.Sprintf to construct and log a message at warn level.
// This call is a wrapper around [Sugaredlogger.Warnf](https://godoc.org/go.uber.org/zap#Sugaredlogger.Warnf)
func Warnf(template string, args ...interface{}) {
	defaultLogger.sLog.Warnf(template, args...)
}

// Warnw logs a message at warn level with some additional context.
// This call is a wrapper around [Sugaredlogger.Warnw](https://godoc.org/go.uber.org/zap#Sugaredlogger.Warnw)
func Warnw(msg string, keysAndValues ...interface{}) {
	defaultLogger.sLog.Warnw(msg, keysAndValues...)
}

// WarnEnabled returns whether output of messages at the warn level is currently enabled.
func WarnEnabled() bool {
	return defaultLogger.Core().Enabled(zap.WarnLevel)
}

// Infoa uses fmt.Sprint to construct and log a message at info level.
// This call is a wrapper around [Sugaredlogger.Info](https://godoc.org/go.uber.org/zap#Sugaredlogger.Info)
func Info(args ...interface{}) {
	defaultLogger.sLog.Info(args...)
}

// Infof uses fmt.Sprintf to construct and log a message at info level.
// This call is a wrapper around [Sugaredlogger.Infof](https://godoc.org/go.uber.org/zap#Sugaredlogger.Infof)
func Infof(template string, args ...interface{}) {
	defaultLogger.sLog.Infof(template, args...)
}

// Infow logs a message at info level with some additional context.
// This call is a wrapper around [Sugaredlogger.Infow](https://godoc.org/go.uber.org/zap#Sugaredlogger.Infow)
func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.sLog.Infow(msg, keysAndValues...)
}

// InfoEnabled returns whether output of messages at the info level is currently enabled.
func InfoEnabled() bool {
	return defaultLogger.Core().Enabled(zap.InfoLevel)
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
// This call is a wrapper around [logger.With](https://godoc.org/go.uber.org/zap#logger.With)
func With(fields ...interface{}) Logger {

	return defaultLogger.With(fields...)
}

// Sync flushes any buffered log entries.
// Processes should normally take care to call Sync before exiting.
// This call is a wrapper around [logger.Sync](https://godoc.org/go.uber.org/zap#logger.Sync)
func Sync() error {
	return defaultLogger.Sync()

}
