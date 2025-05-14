package utils

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const EXECUTABLE_MASK = 0111

// IsExecutableInPath returns the path to a executable file with the provided name.
// The returned pointer is nil if the same was not found. Any errors in reading
// directories & file properties as ignored as the process may not have the
// appropiate permissions for some system directories.
func IsExecutableInPath(cmd string) *string {
	path := os.Getenv("PATH")
	pathDirs := strings.SplitSeq(path, ":")

	for dir := range pathDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == cmd {
				// Check if we have the permission to execute
				fullPath := filepath.Join(dir, entry.Name())
				info, err := os.Stat(fullPath)
				if err != nil {
					continue
				}
				if info.Mode()&EXECUTABLE_MASK != 0 {
					return &fullPath
				}
			}
		}
	}

	return nil
}

// GetAllExecutablesInPath returns a list of all the executable files in the
// directories listed in the PATH environment variable. It ignores any
// directories that cannot be read or do not contain any executable files.
// The returned list does not contain duplicates.
// The name of the executables is returned, not the full path.
func GetAllExecutablesInPath() []string {
	path := os.Getenv("PATH")
	pathDirs := strings.SplitSeq(path, ":")

	executables := make([]string, 0)

	for dir := range pathDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			name := entry.Name()

			// Skip if the entry is a directory or if it is already in the list of executables
			if entry.IsDir() {
				continue
			}
			if slices.Contains(executables, name) {
				continue
			}

			// Check if we have the permission to execute
			fullPath := filepath.Join(dir, name)
			info, err := os.Stat(fullPath)
			if err != nil {
				continue
			}
			if info.Mode()&EXECUTABLE_MASK != 0 {
				executables = append(executables, name)
			}
		}
	}

	return executables
}
