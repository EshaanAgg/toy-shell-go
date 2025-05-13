package shell

import (
	"fmt"
	"os"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

func (s *Shell) ExecuteCommand(line []byte) {
	args, err := utils.GetTokens(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing command: %v\r\n", err)
		return
	}

	// Empty command, no need to execute
	if len(args) == 0 {
		return
	}

	defaultOutput := os.Stdout
	defaultError := os.Stderr

	command := args[0]
	if handler, ok := cmd.HandlerMap[command]; ok {
		handler(args[1:], defaultOutput, defaultError)
	} else {
		cmd.DefaultHandler(args, defaultOutput, defaultError)
	}
}
