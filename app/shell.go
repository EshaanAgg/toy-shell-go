package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EshaanAgg/shell-go/app/cmd"
)

const SHELL_PROMPT = "$ "

func runShell() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(SHELL_PROMPT)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			os.Exit(1)
		}

		input = input[:len(input)-1] // Remove the newline character
		handleCommand(input)
	}
}

func handleCommand(input string) {
	args := strings.Split(input, " ")

	command := args[0]
	leftArgs := args[1:]

	if handler, exists := cmd.HandlerMap[command]; exists {
		handler(leftArgs)
	} else {
		cmd.DefaultHandler(args)
	}
}
