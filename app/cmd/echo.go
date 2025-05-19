package cmd

import (
	"fmt"
	"os"
	"strings"
)

func HandleEcho(args []string, outFile *os.File, errFile *os.File) bool {
	if len(args) == 0 {
		fmt.Fprintf(outFile, "usage: echo <message>, received unexpected args: %v", args)
		return true
	}

	message := strings.Join(args, " ")
	fmt.Fprintf(outFile, "%s", message)
	return true
}
