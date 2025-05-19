package cmd

import (
	"fmt"
	"os"
)

const HOME_DIRECTORY_MARKER = "~"

func HandleCD(args []string, outFile *os.File, errFile *os.File) bool {
	if len(args) != 1 {
		fmt.Fprintf(errFile, "usage: cd <path-to-new-dir>, received unexpected args: %v", args)
		return true
	}

	// Get the path to change the directory to
	path := args[0]
	if path == HOME_DIRECTORY_MARKER {
		path = os.Getenv("HOME")
	}

	// Check if the directory exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Fprintf(errFile, "cd: %s: No such file or directory", path)
		return true
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(errFile, "cd: %s: %v", path, err)
		return true
	}

	return false
}
