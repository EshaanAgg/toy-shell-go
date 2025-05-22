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
