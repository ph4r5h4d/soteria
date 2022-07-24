package zap

import (
	"github.com/ph4r5h4d/soteria/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

type Logger struct {
	logger *zap.Logger
}

func (l Logger) Build() (models.LogInterface, error) {
	logInstance := new(Logger)

	isDebug := false
	envDebug := os.Getenv("SOTERIA_DEBUG")
	if envDebug != "" {
		isDebug, _ = strconv.ParseBool(envDebug)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableCaller = true
	config.DisableStacktrace = true
	if !isDebug {
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	zapLogger, err := config.Build()

	if err != nil {
		return nil, err
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(zapLogger)

	logInstance.logger = zapLogger
	return logInstance, nil
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}
