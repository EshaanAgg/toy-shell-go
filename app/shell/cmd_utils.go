package shell

import (
	"fmt"
	"os"
	"os/exec"

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
