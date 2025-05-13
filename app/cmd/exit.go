package cmd

import (
	"fmt"
	"os"
	"strconv"
)

func HandleExit(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: exit <exit-code>")
	}

	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid exit code: %s", args[0])
		return
	}

	if exitCode < 0 || exitCode > 255 {
		fmt.Printf("Exit code must be between 0 and 255, got: %d", exitCode)
		return
	}

	os.Exit(exitCode)
}
