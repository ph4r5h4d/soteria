package validation

import (
	"os"
)

func fileExistValidator(files *[]string) bool {
	for _, file := range *files {
		if _, err := os.Stat(file); err != nil {
			return false
		}
	}
	return true
}
