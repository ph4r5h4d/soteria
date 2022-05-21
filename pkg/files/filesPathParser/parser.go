package filesPathParser

import (
	"os"
	"path/filepath"
	"strings"
)

func ParseFilesPath(files []string) ([]string, error) {
	var expandedFiles []string

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasPrefix(file, "~/") {
			expandedFiles = append(expandedFiles, filepath.Join(home, file[2:]))
		} else {
			expandedFiles = append(expandedFiles, file)
		}
	}

	return expandedFiles, nil
}
