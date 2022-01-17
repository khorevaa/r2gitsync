package utils

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/khorevaa/logos"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func NewLogger(lg logos.Logger, config gormLogger.Config) gormLogger.Interface {
	return &logger{
		lg:     lg,
		Config: config,
	}
}

type logger struct {
	lg logos.Logger
	gormLogger.Config
}

func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l logger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.lg.Info(fmt.Sprintf(msg, data...), zap.String("filename", utils.FileWithLineNum()))
	}
}

func (l logger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.lg.Warn(fmt.Sprintf(msg, data...), zap.String("filename", utils.FileWithLineNum()))
	}
}

func (l logger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.lg.Error(fmt.Sprintf(msg, data...), zap.String("filename", utils.FileWithLineNum()))
	}
}

func (l logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		strElapced := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
		switch {
		case err != nil && l.LogLevel >= gormLogger.Error:
			var loggerFunc func(msg string, fields ...zap.Field)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				loggerFunc = l.lg.Debug
			} else {
				loggerFunc = l.lg.Error
			}

			sql, rows := fc()
			loggerFunc("Trace SQL",
				zap.String("rows", rowsToString(rows)),
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.String("filename", utils.FileWithLineNum()),
				zap.Error(err))
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLogger.Warn:
			sql, rows := fc()
			l.lg.Warn("Trace SQL",
				zap.String("rows", rowsToString(rows)),
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.String("filename", utils.FileWithLineNum()),
				zap.Duration("slowThreshold", l.SlowThreshold))

		case l.LogLevel >= gormLogger.Info:
			sql, rows := fc()
			l.lg.Info("Trace SQL",
				zap.String("rows", rowsToString(rows)),
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.String("filename", utils.FileWithLineNum()))

		}
	}
}

func rowsToString(rows int64) string {
	if rows == -1 {
		return "-"
	} else {
		return strconv.FormatInt(rows, 10)
	}
}
