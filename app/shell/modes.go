package shell

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func (s *Shell) EnterRAWMode() {
	if s.originalTerminalState != nil {
		panic("Terminal is already in raw mode")
	}

	orgTerm, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Sprintf("Error making terminal raw: %v", err))
	}
	s.originalTerminalState = orgTerm
}

func (s *Shell) ExitRAWMode() {
	if s.originalTerminalState == nil {
		panic("Terminal is not in raw mode")
	}

	if err := term.Restore(int(os.Stdin.Fd()), s.originalTerminalState); err != nil {
		panic(fmt.Sprintf("Error restoring terminal state: %v", err))
	}
	s.originalTerminalState = nil
}
