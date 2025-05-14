package shell

import (
	"strings"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

func (s *Shell) getMatchingCommands() []string {
	allCommands := make([]string, 0)

	// Add all the commands from the cmd package and PATH
	allCommands = append(allCommands, cmd.AllCommands...)
	allCommands = append(allCommands, utils.GetAllExecutablesInPath()...)

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
	s.putChar('\\')
	s.putChar('a')
}

func (s *Shell) handleTabClick() {
	matches := s.getMatchingCommands()

	if len(matches) == 0 {
		s.printBell()
		return
	}

	// Handle only the first match
	matchedCmd := matches[0]
	// Print the leftover part of the command
	for i := len(s.input); i < len(matchedCmd); i++ {
		s.putChar(matchedCmd[i])
	}
}
