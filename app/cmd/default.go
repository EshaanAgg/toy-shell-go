package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/EshaanAgg/shell-go/app/utils"
)

func executeBinary(cmd string, args ...string) {
	p := exec.Command(cmd, args...)
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr
	err := p.Run()
	if err != nil {
		fmt.Printf("the command exited with error: %v", err)
	}
}

func DefaultHandler(args []string) {
	if len(args) == 0 {
		fmt.Println()
		return
	}

	// Check for executable in path
	cmd := args[0]
	path, err := utils.IsExecutableInPath(cmd)
	if err != nil {
		fmt.Printf("error while parsing the folders in PATH: %v", err)
		return
	}
	// Run the exectuable
	if path != nil {
		executeBinary(cmd, args[1:]...)
		return
	}

	// Unrecognized
	fmt.Printf("%s: command not found\n", args[0])
}
