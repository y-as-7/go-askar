package debugger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	Debugger *Debugger
	Config   logger.Config
}

func NewGormLogger(d *Debugger) *GormLogger {
	return &GormLogger{
		Debugger: d,
		Config: logger.Config{
			LogLevel: logger.Info,
		},
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.Config.LogLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Debugger == nil {
		return
	}
	sql, rows := fc()
	duration := time.Since(begin)
	_ = rows // we could store this too

	l.Debugger.AddQuery(sql, duration)
}
