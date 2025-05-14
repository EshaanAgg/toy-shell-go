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
	// Get a unique list of all commands from cmd package and executables in PATH
	allCommands := make([]string, 0)
	allCommands = append(allCommands, utils.GetAllExecutablesInPath()...)
	for _, cmd := range cmd.AllCommands {
		if !slices.Contains(allCommands, cmd) {
			allCommands = append(allCommands, cmd)
		}
	}

	// Filter out the commands that match the current input
	matchedCommands := make([]string, 0)
	curInput := string(s.input)
	for _, cmd := range allCommands {
		if cmd != curInput && strings.HasPrefix(cmd, curInput) {
			matchedCommands = append(matchedCommands, cmd)
		}
	}

	if len(matchedCommands) == 0 {
		return matchedCommands
	}

	// Sort the matched commands and see if they are all
	// the prefix of each other. Then we should only return the
	// first one.
	slices.Sort(matchedCommands)
	allPrefix := true
	for i := 1; i < len(matchedCommands); i++ {
		if !strings.HasPrefix(matchedCommands[i], matchedCommands[i-1]) {
			allPrefix = false
			break
		}
	}

	if allPrefix {
		return []string{matchedCommands[0]}
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

func (s *Shell) handleMultipleMatches(matches []string) {
	// On first multiple match, ring the bell
	if !s.hadMultipleMatched {
		s.hadMultipleMatched = true
		s.printBell()
		return
	}

	s.hadMultipleMatched = false

	// Go to next line and print the matches
	fmt.Fprintf(os.Stdout, "\r\n")
	for _, match := range matches {
		fmt.Fprintf(os.Stdout, "%s ", match)
	}
	// Go to the next line, and print the prompt & the input
	fmt.Fprintf(os.Stdout, "\r\n")
	s.putPrompt()
	fmt.Fprintf(os.Stdout, "%s", string(s.input))
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

	s.handleMultipleMatches(matches)
}
