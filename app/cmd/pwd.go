package cmd

import (
	"fmt"
	"os"
)

func HandlePWD(args []string, outFile *os.File, errFile *os.File) {
	if len(args) != 0 {
		fmt.Fprintf(errFile, "usage: pwd, received unexpected args: %v", args)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(errFile, "unable to get the current working directory: %v", err)
		return
	}
	fmt.Fprintf(outFile, "%s\r\n", pwd)
}
