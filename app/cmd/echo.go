package cmd

import "fmt"

func HandleEcho(args []string) {
	l := len(args)

	if l == 0 {
		return
	}

	for i, str := range args {
		if i == l-1 {
			fmt.Print(str)
		} else {
			fmt.Printf("%s ", str)
		}
	}
}
