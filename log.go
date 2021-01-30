package liblog

import (
	"context"

	"github.com/jiuzhou-zhao/go-fundamental/interfaces"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() (interfaces.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		logger: logger,
	}, nil
}

func (logger *ZapLogger) Record(ctx context.Context, depth int, level interfaces.LoggerLevel, v ...interface{}) {
	log := logger.logger.WithOptions(zap.AddCallerSkip(depth + 1)).Sugar()
	switch level {
	case interfaces.LogLevelDebug:
		log.Debug(v...)
	case interfaces.LogLevelInfo:
		log.Info(v...)
	case interfaces.LogLevelWarn:
		log.Warn(v...)
	case interfaces.LogLevelError:
		log.Error(v...)
	case interfaces.LogLevelFatal:
		log.Fatal(v...)
	}
}

func (logger *ZapLogger) Recordf(ctx context.Context, depth int, level interfaces.LoggerLevel, format string, v ...interface{}) {
	log := logger.logger.WithOptions(zap.AddCallerSkip(depth + 1)).Sugar()
	switch level {
	case interfaces.LogLevelDebug:
		log.Debugf(format, v...)
	case interfaces.LogLevelInfo:
		log.Infof(format, v...)
	case interfaces.LogLevelWarn:
		log.Warnf(format, v...)
	case interfaces.LogLevelError:
		log.Errorf(format, v...)
	case interfaces.LogLevelFatal:
		log.Fatalf(format, v...)
	}
}
