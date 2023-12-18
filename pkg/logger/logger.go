package logger

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func New() error {
	log, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zapLogger = log
	return nil
}

func Debug(msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}
