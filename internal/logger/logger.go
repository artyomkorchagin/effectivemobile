package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type devMode bool

const (
	Production  devMode = false
	Development devMode = true
)

func New(mode devMode) (*zap.Logger, error) {
	if mode == Development {
		return newDevelopmentLogger()
	}
	return newProductionLogger()
}

func newProductionLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	return config.Build()
}

func newDevelopmentLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return config.Build()
}
