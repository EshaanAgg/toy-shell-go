package cmd

import (
	"fmt"
	"os"
)

func HandleCD(args []string) {
	if len(args) != 1 {
		fmt.Printf("usage: cd <path-to-new-dir>, recieved unexpected args: %v\n", args)
		return
	}

	err := os.Chdir(args[0])
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
		return
	}
}
