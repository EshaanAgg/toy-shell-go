package main

import (
	"fmt"
	"os"

	"github.com/EshaanAgg/shell-go/app/parser"
	"github.com/EshaanAgg/shell-go/app/shell"
	"golang.org/x/term"
)

func main() {
	// Put the terminal in raw mode
	// This is necessary to read the input character by character
	oldTerminalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Sprintf("Error making terminal raw: %v", err))
	}
	defer term.Restore(int(os.Stdin.Fd()), oldTerminalState)

	// Register the interrupt handler
	// Create the REPL loop with the parser and shell
	p := parser.NewParser()
	s := shell.NewShell()

	for {
		args, err := p.GetCommand(s)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if args == nil {
			// User pressed Ctrl+C, exit the loop
			break
		}
		s.ExecuteCommand(args)
	}
}
