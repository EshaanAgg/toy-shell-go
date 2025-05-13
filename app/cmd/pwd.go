package cmd

import (
	"fmt"
	"io"
	"os"
)

func HandlePWD(args []string, outFile io.Writer, errFile io.Writer) {
	if len(args) != 0 {
		fmt.Fprintf(errFile, "usage: pwd, received unexpected args: %v\r\n", args)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(errFile, "unable to get the current working directory: %v\r\n", err)
		return
	}
	fmt.Fprintf(outFile, "%s\r\n", pwd)
}
