package cmd

import (
	"fmt"
	"slices"

	"github.com/EshaanAgg/shell-go/app/utils"
)

func HandleType(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: type <command>")
		return
	}

	cmd := args[0]

	// Check for shell built-in
	exists := slices.Index(AllCommands, cmd)
	if exists != -1 {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	// Check for executable in path
	path := utils.IsExecutableInPath(cmd)
	if path != nil {
		fmt.Printf("%s is %s\n", cmd, *path)
		return
	}

	// Unrecognized
	fmt.Printf("%s: not found\n", cmd)
}
