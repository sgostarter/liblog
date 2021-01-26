package liblog

import (
	"context"

	"github.com/jiuzhou-zhao/go-fundamental/interfaces"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (interfaces.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		logger: logger.Sugar(),
	}, nil
}

func (logger *ZapLogger) Record(ctx context.Context, level interfaces.LoggerLevel, v ...interface{}) {
	switch level {
	case interfaces.LogLevelDebug:
		logger.logger.Debug(v...)
	case interfaces.LogLevelInfo:
		logger.logger.Info(v...)
	case interfaces.LogLevelWarn:
		logger.logger.Warn(v...)
	case interfaces.LogLevelError:
		logger.logger.Error(v...)
	case interfaces.LogLevelFatal:
		logger.logger.Fatal(v...)
	}
}

func (logger *ZapLogger) Recordf(ctx context.Context, level interfaces.LoggerLevel, format string, v ...interface{}) {
	switch level {
	case interfaces.LogLevelDebug:
		logger.logger.Debugf(format, v...)
	case interfaces.LogLevelInfo:
		logger.logger.Infof(format, v...)
	case interfaces.LogLevelWarn:
		logger.logger.Warnf(format, v...)
	case interfaces.LogLevelError:
		logger.logger.Errorf(format, v...)
	case interfaces.LogLevelFatal:
		logger.logger.Fatalf(format, v...)
	}
}
