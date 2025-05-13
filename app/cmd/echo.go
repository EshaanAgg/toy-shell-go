package cmd

import (
	"fmt"
	"io"
	"strings"
)

func HandleEcho(args []string, outFile io.Writer, errFile io.Writer) {
	if len(args) == 0 {
		fmt.Fprintf(outFile, "usage: echo <message>, received unexpected args: %v\r\n", args)
		return
	}

	message := strings.Join(args, " ")
	fmt.Fprintf(outFile, "%s\r\n", message)
}
