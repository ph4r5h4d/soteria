package validation

import (
	"go.uber.org/zap"
)

var validators = map[string]func(files *[]string) bool{
	"fileExists": fileExistValidator,
}

func Validate(files []string, logger *zap.Logger) bool {
	logger.Info("Validating configuration")
	for i, v := range validators {
		logger.Debug("Running <" + i + "> validation...")
		v(&files)
		logger.Debug("<" + i + "> validation completed.")
	}
	logger.Info("Validation completed.")
	return true
}
