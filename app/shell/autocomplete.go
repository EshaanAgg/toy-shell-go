package shell

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

func (s *Shell) getMatchingCommands() []string {
	allCommands := make([]string, 0)

	// Add all the commands from the cmd package and PATH
	allCommands = append(allCommands, utils.GetAllExecutablesInPath()...)
	for _, cmd := range cmd.AllCommands {
		if !slices.Contains(allCommands, cmd) {
			allCommands = append(allCommands, cmd)
		}
	}

	matchedCommands := make([]string, 0)
	curInput := string(s.input)
	for _, cmd := range allCommands {
		if cmd != curInput && strings.HasPrefix(cmd, curInput) {
			matchedCommands = append(matchedCommands, cmd)
		}
	}

	return matchedCommands
}

func (s *Shell) printBell() {
	fmt.Fprintf(os.Stdout, "\a")
}

func (s *Shell) handleOneMatch(cmd string) {
	// Print the leftover part of the command
	for i := len(s.input); i < len(cmd); i++ {
		s.putChar(cmd[i])
	}
	s.putChar(' ')
}

func (s *Shell) handleTabClick() {
	matches := s.getMatchingCommands()

	if len(matches) == 0 {
		s.printBell()
		return
	}

	if len(matches) == 1 {
		s.handleOneMatch(matches[0])
	}
}
