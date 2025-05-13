package shell

import (
	"fmt"
	"io"
	"os"
)

// checkForRedirection checks if the command has any redirection operators.
// If it does, it opens the appropriate file for writing and
// sets the outFile or errFile accordingly.
// Currently only single redirection, provided at the end of the command is supported.
func (c *command) checkForRedirection() error {
	// For direction, there should be atleast 3 args:
	// <prev-command> <redirection-op> <file>
	l := len(c.args)
	if l < 3 {
		return nil
	}

	// Check for redirection
	fileName := c.args[l-1]
	redirected := false

	switch c.args[l-2] {
	case ">", "1>":
		f, err := getFile(fileName, false)
		if err != nil {
			return fmt.Errorf("unable to open file %s for redirecting STDOUT: %v", fileName, err)
		}
		c.outFile = f
		redirected = true

	case "2>":
		f, err := getFile(fileName, false)
		if err != nil {
			return fmt.Errorf("unable to open file %s for redirecting STDERR: %v", fileName, err)
		}
		c.errFile = f
		redirected = true

	case ">>", "1>>":
		f, err := getFile(fileName, true)
		if err != nil {
			return fmt.Errorf("unable to open file %s for appending STDOUT: %v", fileName, err)
		}
		c.outFile = f
		redirected = true

	case "2>>":
		f, err := getFile(c.args[l-1], true)
		if err != nil {
			return fmt.Errorf("unable to open file %s for appending STDERR: %v", c.args[l-1], err)
		}
		c.errFile = f
		redirected = true
	}

	// If the command is redirected, remove the last two args
	// <redirection-op> and <file>
	if redirected {
		c.args = c.args[:l-2]
	}

	return nil
}

// getFile opens a file for writing. If append is true, it opens the file in append mode.
// Otherwise, it creates a new file. If the file already exists, it truncates it.
func getFile(fileName string, append bool) (io.Writer, error) {
	if append {
		appendMode := os.O_APPEND | os.O_CREATE
		return os.OpenFile(fileName, appendMode, 0644)
	}

	return os.Create(fileName)
}
