package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/EshaanAgg/shell-go/app/utils"
)

func DefaultHandler(args []string, outFile *os.File, errFile *os.File) {
	// Check for executable in path
	cmd := args[0]
	path := utils.IsExecutableInPath(cmd)

	// Run the exectuable
	if path != nil {
		p := exec.Command(*path, args[1:]...)
		p.Stdout = outFile
		p.Stderr = errFile
		if err := p.Run(); err != nil {
			fmt.Fprintf(errFile, "the command exited with error: %v\r\n", err)
		}
		return
	}

	// Unrecognized
	fmt.Fprintf(errFile, "%s: command not found\r\n", cmd)
}
