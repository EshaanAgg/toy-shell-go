package parser

import (
	"fmt"
	"os"
	"strings"
)

func (p *Parser) putCLRF() {
	fmt.Print(CLRF)
}

func (p *Parser) getTokens() []string {
	s := string(p.currentInput)
	return strings.Split(s, " ")
}

func (p *Parser) readByte() byte {
	b := make([]byte, 1)
	if _, err := os.Stdin.Read(b); err != nil {
		panic(fmt.Sprintf("Error reading byte: %v", err))
	}
	return b[0]
}
