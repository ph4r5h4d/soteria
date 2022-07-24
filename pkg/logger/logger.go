package logger

import (
	"errors"
	"github.com/ph4r5h4d/soteria/models"
	"github.com/ph4r5h4d/soteria/pkg/logger/zap"
)

var loggers = map[string]interface{}{
	"zap": &zap.Logger{},
}

func BuildLogger(name string) (models.LogInterface, error) {
	loggerInstance, ok := loggers[name]
	if !ok {
		return nil, errors.New("unsupported logger")
	}

	logger, err := loggerInstance.(models.LogInterface).Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
