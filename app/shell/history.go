package shell

import (
	"fmt"
	"strconv"
)

func (c *command) handleHistory(s *Shell) {
	startIdx := 0
	if len(c.args) > 1 {
		cmdCnt, err := strconv.Atoi(c.args[1])
		if err != nil || cmdCnt < 0 {
			fmt.Fprintf(c.errFile, "Invalid argument: %s\r\n", c.args[1])
			return
		}

		startIdx = len(s.history) - cmdCnt
	}

	for i := startIdx; i < len(s.history); i++ {
		fmt.Fprintf(c.outFile, "\t%d %s\r\n", i+1, s.history[i])
	}
}

func (s *Shell) clearLine() {
	fmt.Printf("\r%s", MOVE_CURSOR_LEFT)
	fmt.Print(CLEAR_LINE)
	fmt.Print(MOVE_CURSOR_LEFT)
}

func (s *Shell) putFromHistory() {
	if s.curHistoryIdx == -1 || s.curHistoryIdx >= len(s.history) {
		return
	}

	s.clearLine()
	s.input = []byte(s.history[s.curHistoryIdx])
	s.cursorPosition = len(s.input)
	fmt.Printf("$ %s", s.input)
}

func (s *Shell) handleUpArrowPress() {
	if s.curHistoryIdx == -1 {
		s.curHistoryIdx = len(s.history) - 1
	} else if s.curHistoryIdx > 0 {
		s.curHistoryIdx--
	}
	s.putFromHistory()
}

func (s *Shell) hanldeDownArrowPress() {
	if s.curHistoryIdx == -1 {
		return
	}

	if s.curHistoryIdx < len(s.history)-1 {
		s.curHistoryIdx++
	} else {
		s.curHistoryIdx = -1
	}

	s.putFromHistory()
}
