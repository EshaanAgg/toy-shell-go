package cmd

import "fmt"

func HandleEcho(args []string) {
	l := len(args)

	if l == 0 {
		return
	}

	for i, str := range args {
		if i == l-1 {
			// If it's the last string, print a newline after it
			fmt.Printf("%s\n", str)
		} else {
			// Find a space after the intermediate strings
			fmt.Printf("%s ", str)
		}
	}
}
