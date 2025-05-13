package cmd

import (
	"fmt"
	"os"
)

func HandlePWD(args []string) {
	if len(args) != 0 {
		fmt.Printf("usage: pwd, recieved unexpected args: %v\n", args)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("unable to get the current working directory: %v", err)
	}
	fmt.Println(pwd)
}
