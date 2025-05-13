package shell

import (
	"fmt"
)

type Shell struct {
	shellPrompt string
}

func NewShell() *Shell {
	return &Shell{
		shellPrompt: "$ ",
	}
}

func (s *Shell) PutPrompt() {
	fmt.Print(s.shellPrompt)
}
