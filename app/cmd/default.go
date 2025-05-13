package cmd

import "fmt"

func DefaultHandler(args []string) {
	fmt.Printf("%s: command not found\n", args[0])
}
