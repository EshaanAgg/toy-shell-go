package cmd

import (
	"fmt"
	"io"
	"slices"

	"github.com/EshaanAgg/shell-go/app/utils"
)

func HandleType(args []string, outFile io.Writer, errFile io.Writer) {
	if len(args) == 0 {
		fmt.Fprintf(outFile, "usage: type <command>, received unexpected args: %v\r\n", args)
		return
	}

	cmd := args[0]

	// Check for shell built-in
	exists := slices.Index(AllCommands, cmd)
	if exists != -1 {
		fmt.Fprintf(outFile, "%s is a shell builtin\r\n", cmd)
		return
	}

	// Check for executable in path
	path := utils.IsExecutableInPath(cmd)
	if path != nil {
		fmt.Fprintf(outFile, "%s is %s\r\n", cmd, *path)
		return
	}

	// Unrecognized
	fmt.Fprintf(errFile, "%s: not found\r\n", cmd)
}
