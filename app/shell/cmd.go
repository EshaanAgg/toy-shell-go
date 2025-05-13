package shell

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/EshaanAgg/shell-go/app/cmd"
	"github.com/EshaanAgg/shell-go/app/utils"
)

// defaultCommandHandler is the default command handler for unrecognized commands.
// This is defined here and not in the "cmd" package because it needs access to the
// shell's environment and state.
func (s *Shell) defaultCommandHandler(args []string, outFile *os.File, errFile *os.File) {
	// Check for executable in path
	cmd := args[0]
	path := utils.IsExecutableInPath(cmd)

	// Run the exectuable
	if path != nil {
		s.ExitRAWMode()

		p := exec.Command(cmd, args[1:]...)
		p.Stdout = outFile
		p.Stderr = errFile
		p.Stdin = os.Stdin
		p.Run()

		s.EnterRAWMode()
		return
	}

	// Unrecognized
	fmt.Fprintf(errFile, "%s: command not found\r\n", cmd)
}

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
		s.defaultCommandHandler(args, defaultOutput, defaultError)
	}
}
