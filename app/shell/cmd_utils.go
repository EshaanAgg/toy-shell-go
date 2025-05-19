package shell

import (
	"fmt"
	"os"
	"os/exec"
	"slices"

	"github.com/EshaanAgg/shell-go/app/utils"
)

var standardOSFiles = []*os.File{
	os.Stdin,
	os.Stdout,
	os.Stderr,
}

// Represents a single command that can be executed
// in the shell. It should have no redirection or piping.
type command struct {
	args    []string
	inFile  *os.File
	outFile *os.File
	errFile *os.File
}

// Executes the command in using a process on the OS.
func (c *command) executeOnOS() {
	cmd := c.args[0]

	// Check if the command is in the PATH
	path := utils.IsExecutableInPath(cmd)
	if path == nil {
		fmt.Fprintf(c.errFile, "%s: command not found\r\n", cmd)
		return
	}

	p := exec.Command(c.args[0], c.args[1:]...)
	p.Stdout = c.outFile
	p.Stderr = c.errFile
	p.Stdin = c.inFile

	p.Run()
	c.cleanup()
}

func (c *command) cleanup() {
	if !slices.Contains(standardOSFiles, c.inFile) {
		c.inFile.Close()
	}
	if !slices.Contains(standardOSFiles, c.outFile) {
		c.outFile.Close()
	}
	if !slices.Contains(standardOSFiles, c.errFile) {
		c.errFile.Close()
	}
}
