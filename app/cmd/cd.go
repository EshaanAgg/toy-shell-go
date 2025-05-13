package cmd

import (
	"fmt"
	"os"
)

const HOME_DIRECTORY_MARKER = "~"

func HandleCD(args []string, outFile *os.File, errFile *os.File) {
	if len(args) != 1 {
		fmt.Fprintf(errFile, "usage: cd <path-to-new-dir>, received unexpected args: %v\r\n", args)
		return
	}

	// Get the path to change the directory to
	path := args[0]
	if path == HOME_DIRECTORY_MARKER {
		path = os.Getenv("HOME")
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(errFile, "%v\r\n", err)
		return
	}
}
