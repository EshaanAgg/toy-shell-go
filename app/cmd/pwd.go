package cmd

import (
	"fmt"
	"os"
)

func HandlePWD(args []string, outFile *os.File, errFile *os.File) bool {
	if len(args) != 0 {
		fmt.Fprintf(errFile, "usage: pwd, received unexpected args: %v", args)
		return true
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(errFile, "unable to get the current working directory: %v", err)
		return true
	}

	fmt.Fprintf(outFile, "%s", pwd)
	return true
}
