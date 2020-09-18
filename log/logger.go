package log

import (
	"go.uber.org/zap"
)

var _ Logger = (*logger)(nil)

//var _ Logger = (*zap.SugaredLogger)(nil)

type logger struct {
	*zap.Logger
	sLog *zap.SugaredLogger
}

func (log *logger) Named(name string) Logger {

	nLog := log.Logger.Named(name)
	return &logger{
		Logger: nLog,
		sLog:   nLog.Sugar(),
	}
}

func (log *logger) Debug(args ...interface{}) {
	log.sLog.Debug(args...)
}
func (log *logger) Debuga(msg string, fields ...zap.Field) {
	log.Logger.Debug(msg, fields...)
}

func (log *logger) Debugf(template string, args ...interface{}) {
	log.sLog.Debugf(template, args...)
}

func (log *logger) Debugw(msg string, keysAndValues ...interface{}) {
	log.sLog.Debugw(msg, keysAndValues...)
}

func (log *logger) DebugEnabled() bool {
	return log.Core().Enabled(zap.DebugLevel)
}

func (log *logger) Error(args ...interface{}) {
	log.sLog.Error(args...)
}

func (log *logger) Errorf(template string, args ...interface{}) {
	log.sLog.Errorf(template, args...)
}

func (log *logger) Errorw(msg string, keysAndValues ...interface{}) {
	log.sLog.Errorw(msg, keysAndValues...)
}

func (log *logger) Warn(args ...interface{}) {
	log.sLog.Warn(args...)
}

func (log *logger) Warnf(template string, args ...interface{}) {
	log.sLog.Warnf(template, args...)
}

func (log *logger) Warnw(msg string, keysAndValues ...interface{}) {
	log.sLog.Warnw(msg, keysAndValues...)
}

func (log *logger) Info(args ...interface{}) {
	log.sLog.Info(args...)
}

func (log *logger) Infof(template string, args ...interface{}) {
	log.sLog.Infof(template, args...)
}

func (log *logger) Infow(msg string, keysAndValues ...interface{}) {
	log.sLog.Infow(msg, keysAndValues...)
}

func (log *logger) With(args ...interface{}) Logger {

	sLog := log.sLog.With(args...)

	newLog := &logger{
		Logger: sLog.Desugar(),
		sLog:   sLog,
	}

	return newLog
}

var defaultLogger = newLogger()

func newLogger() *logger {

	cfg := zap.NewProductionConfig()
	log, _ := cfg.Build(zap.AddCallerSkip(1))
	sugar := log.Sugar()

	return &logger{
		Logger: log,
		sLog:   sugar,
	}
}

func NewLogger() Logger {
	return newLogger()
}

func (log *logger) SetDebug() {

	log.Logger, _ = zap.NewDevelopment(zap.AddCallerSkip(1))
	log.sLog = log.Sugar()

}
