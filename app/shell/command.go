package shell

import (
	"fmt"
	"os"
	"os/exec"
	"slices"

	"github.com/EshaanAgg/shell-go/app/cmd"
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

// newCommand creates a new command from the given line.
// It parses the line into arguments and checks for
// redirection. If redirection is found, it sets the
// appropriate output and error files. Returns an error if
// the command is invalid or if redirection fails.
func newCommand(line []byte) (*command, error) {
	args, err := utils.GetTokens(line)
	if err != nil {
		return nil, fmt.Errorf("error parsing command: %v", err)
	}

	cmd := &command{
		args:    args,
		outFile: os.Stdout,
		errFile: os.Stderr,
		inFile:  os.Stdin,
	}
	if err := cmd.checkForRedirection(); err != nil {
		return nil, fmt.Errorf("redirection error: %v", err)
	}

	return cmd, nil
}

// execute runs the command on the shell. If the command is an inbuilt
// command, it will be executed using the handler map. If the command
// is not found in the handler map, it will be executed using the OS.
func (c *command) execute(s *Shell, inPipeline bool) {
	// Empty command, no need to execute
	if len(c.args) == 0 {
		return
	}

	command := c.args[0]
	if handler, ok := cmd.HandlerMap[command]; ok {
		// Found a handler registered for the command
		hadOutput := handler(c.args[1:], c.outFile, c.errFile)

		if inPipeline {
			// If we are in a pipeline, then the shell is not in RAW mode
			// We can print a new line before the prompt
			fmt.Fprint(c.outFile, "\n")
		} else if hadOutput {
			// We must move the cursor to the next line and reset it's position
			fmt.Fprintf(c.outFile, "\r\n")
		}
		c.cleanup()
	} else {
		if !inPipeline {
			s.ExitRAWMode()
			defer s.EnterRAWMode()
		}
		c.executeOnOS()
	}
}

// Executes the command in using a process on the OS.
// The shell is assumed NOT to be in RAW mode when this is called.
func (c *command) executeOnOS() {
	cmd := c.args[0]

	// Check if the command is in the PATH
	path := utils.IsExecutableInPath(cmd)
	if path == nil {
		fmt.Fprintf(c.errFile, "%s: command not found\n", cmd)
		return
	}

	p := exec.Command(c.args[0], c.args[1:]...)
	p.Stdin = c.inFile
	p.Stdout = c.outFile
	p.Stderr = c.errFile

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
