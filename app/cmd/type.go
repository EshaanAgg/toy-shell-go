package cmd

import (
	"fmt"
	"os"
	"slices"

	"github.com/EshaanAgg/shell-go/app/utils"
)

func HandleType(args []string, outFile *os.File, errFile *os.File) {
	if len(args) == 0 {
		fmt.Fprintf(outFile, "usage: type <command>, received unexpected args: %v", args)
		return
	}

	cmd := args[0]

	// Check for shell built-in
	exists := slices.Index(AllCommands, cmd)
	if exists != -1 {
		fmt.Fprintf(outFile, "%s is a shell builtin", cmd)
		return
	}

	// Check for executable in path
	path := utils.IsExecutableInPath(cmd)
	if path != nil {
		fmt.Fprintf(outFile, "%s is %s", cmd, *path)
		return
	}

	// Unrecognized
	fmt.Fprintf(errFile, "%s: not found", cmd)
}
