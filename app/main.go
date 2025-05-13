package main

import (
	"github.com/EshaanAgg/shell-go/app/parser"
	"github.com/EshaanAgg/shell-go/app/shell"
)

func main() {
	p := parser.NewParser()
	s := shell.NewShell()

	for {
		// Run an infinite loop of reading commands
		// and executing them
		p.HandleCommand(s)
	}
}
