package shell

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type pipelineCommand struct {
	commands []*command
}

func newPipelineCommand(line []byte) (*pipelineCommand, error) {
	parts := bytes.Split(line, []byte("|"))

	// Create commands
	commands := make([]*command, 0, len(parts))
	for i, part := range parts {
		cmd, err := newCommand(part)
		if err != nil {
			return nil, fmt.Errorf("error creating command[%d]: %v", i, err)
		}
		commands = append(commands, cmd)
	}

	// Map the output of one command to the input of the next
	for idx := range len(commands) - 1 {
		r, w, err := os.Pipe()
		if err != nil {
			return nil, fmt.Errorf("error creating pipe: %v", err)
		}
		commands[idx].outFile = w
		commands[idx+1].inFile = r
	}

	return &pipelineCommand{
		commands: commands,
	}, nil
}

func (p *pipelineCommand) execute(s *Shell) {
	wg := sync.WaitGroup{}
	wg.Add(len(p.commands))

	s.ExitRAWMode()

	// Execute all commands in parallel
	for _, cmd := range p.commands {
		go func(c *command) {
			defer wg.Done()
			c.execute(s, false)
		}(cmd)
	}

	wg.Wait()
	s.EnterRAWMode()
}
