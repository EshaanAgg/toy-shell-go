package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const EXECUTABLE_MASK = 0111

func IsExecutableInPath(cmd string) (*string, error) {
	path := os.Getenv("PATH")
	pathDirs := strings.SplitSeq(path, ":")

	for dir := range pathDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("unable to read entries in dir '%s': %w", dir, err)
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == cmd {
				// Check if we have the permission to execute
				fullPath := filepath.Join(dir, entry.Name())
				info, err := os.Stat(fullPath)
				if err != nil {
					return nil, fmt.Errorf("unable to get the permissions for the executable '%s': %w", fullPath, err)
				}
				if info.Mode()&EXECUTABLE_MASK != 0 {
					return &fullPath, nil
				}
			}
		}
	}

	return nil, nil
}
