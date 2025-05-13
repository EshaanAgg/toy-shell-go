package cmd

import (
	"fmt"
	"os"
)

const HOME_DIRECTORY_MARKER = "~"

func HandleCD(args []string) {
	if len(args) != 1 {
		fmt.Printf("usage: cd <path-to-new-dir>, recieved unexpected args: %v\n", args)
		return
	}

	// Get the path to change the directory to
	path := args[0]
	if path == HOME_DIRECTORY_MARKER {
		path = os.Getenv("HOME")
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
		return
	}
}
