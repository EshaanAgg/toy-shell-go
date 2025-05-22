package shell

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

type Shell struct {
	originalTerminalState *term.State

	input          []byte
	cursorPosition int
	history        []string

	hadMultipleMatched bool
}

func NewShell() *Shell {
	s := &Shell{}
	s.EnterRAWMode()
	return s
}

func (s *Shell) Kill() {
	s.ExitRAWMode()
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
	fmt.Print("$ ")
}

func (s *Shell) ExecuteCommand(line []byte) {
	if bytes.Contains(line, []byte("|")) {
		s.ExecutePipelineCommand(line)
		return
	}

	cmd, err := newCommand(line)
	if err != nil {
		fmt.Printf("Error creating command: %v\r\n", err)
		return
	}
	cmd.execute(s, false)
}

func (s *Shell) ExecutePipelineCommand(line []byte) {
	pipelineCmd, err := newPipelineCommand(line)
	if err != nil {
		fmt.Printf("Error creating pipeline command: %v\r\n", err)
		return
	}
	pipelineCmd.execute(s)
}
