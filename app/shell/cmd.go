package shell

import (
	"fmt"
	"os"

	"github.com/EshaanAgg/shell-go/app/cmd"
)

func (s *Shell) ExecuteCommand(args []string) {
	// Print a new line before executing the command
	// This is so that the output of the command is not printed on the same line as the prompt
	fmt.Print("\n")

	if len(args) == 0 {
		return
	}

	command := args[0]
	leftArgs := args[1:]

	if handler, exists := cmd.HandlerMap[command]; exists {
		handler(leftArgs, os.Stdout, os.Stderr)
	} else {
		cmd.DefaultHandler(args, os.Stdout, os.Stderr)
	}
}
