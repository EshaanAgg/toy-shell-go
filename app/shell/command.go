package shell

import (
	"fmt"
	"os"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

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

func (c *command) execute(s *Shell) {
	// Empty command, no need to execute
	if len(c.args) == 0 {
		return
	}

	command := c.args[0]
	if handler, ok := cmd.HandlerMap[command]; ok {
		// Found a handler registered for the command
		handler(c.args[1:], c.outFile, c.errFile)
		c.cleanup()
	} else {
		// Execute the command on the OS after entering raw mode
		s.ExitRAWMode()
		c.executeOnOS()
		s.EnterRAWMode()
	}
}
