package cmd

import (
	"fmt"
	"os"
	"strings"
)

func HandleEcho(args []string, outFile *os.File, errFile *os.File) {
	defer cleanup(outFile, errFile)

	if len(args) == 0 {
		fmt.Fprintf(outFile, "usage: echo <message>, received unexpected args: %v\r\n", args)
		return
	}

	message := strings.Join(args, " ")
	fmt.Fprintf(outFile, "%s\r\n", message)
}
