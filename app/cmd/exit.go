package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func HandleExit(args []string, outFile io.Writer, errFile io.Writer) {
	if len(args) != 1 {
		fmt.Fprintf(errFile, "usage: exit <exit-code>, received unexpected args: %v\r\n", args)
	}

	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(errFile, "invalid exit code: %s\r\n", args[0])
		return
	}

	if exitCode < 0 || exitCode > 255 {
		fmt.Fprintf(errFile, "exit code must be between 0 and 255, got: %d\r\n", exitCode)
		return
	}

	os.Exit(exitCode)
}
