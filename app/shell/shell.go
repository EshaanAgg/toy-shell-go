package shell

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Shell struct {
	originalTerminalState *term.State

	input          []byte
	cursorPosition int
}

func NewShell() *Shell {
	oldTerminalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Sprintf("Error making terminal raw: %v", err))
	}
	return &Shell{
		originalTerminalState: oldTerminalState,
	}
}

func (s *Shell) Kill() {
	// Restore the original terminal state
	if err := term.Restore(int(os.Stdin.Fd()), s.originalTerminalState); err != nil {
		panic(fmt.Sprintf("Error restoring terminal state: %v", err))
	}
	os.Exit(0)
}

func (s *Shell) readByte() byte {
	b := make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		panic("Error reading byte from stdin")
	}
	return b[0]
}

func (s *Shell) putPrompt() {
	// Always move the cursor to the beginning of the line
	// before printing the prompt
	fmt.Print("\r$ ")
}
