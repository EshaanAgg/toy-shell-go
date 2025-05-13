package shell

import (
	"fmt"
	"os"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

// Represents a single command that can be executed
// in the shell. It should have no redirection or piping.
type command struct {
	args    []string
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
		handler(c.args[1:], c.outFile, c.errFile)
	} else {
		s.defaultCommandHandler(c.args, c.outFile, c.errFile)
	}
}
