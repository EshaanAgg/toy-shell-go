package cmd

import (
	"fmt"
	"slices"
)

func HandleType(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: type <command>")
		return
	}

	cmd := args[0]
	exists := slices.Index(AllCommands, cmd)
	if exists != -1 {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else {
		fmt.Printf("%s: not found\n", cmd)
	}
}
