package parser

import (
	"fmt"

	"github.com/EshaanAgg/shell-go/app/shell"
)

type Parser struct {
	// Information about the current token being parsed
	currentInput []byte
	position     int

	// Shows if the parser is currently in a quoted string.
	// If nil, then the parser is not in a quoted string.
	// If not nil, then the parser is in a quoted string and this byte
	// is the quote character.
	quoted *byte
}

func NewParser() *Parser {
	return &Parser{
		currentInput: make([]byte, 0),
		position:     0,
		quoted:       nil,
	}
}

// Reads a complete command from the STDIN character by character,
// and when completely parsed, returns the same to the caller.
func (p *Parser) GetCommand(s *shell.Shell) ([]string, error) {
	s.PutPrompt()

	// Read the input from the user character by character
	for {
		char := p.readByte()

		switch char {
		case KEY_CTRL_C:
			fmt.Print("\n\rExiting shell...\r\n")
			return nil, nil

		case KEY_ENTER, KEY_NEWLINE:
			p.putCLRF()
			toks := p.getTokens()
			return toks, nil

		default:
			fmt.Print(string(char))
			p.currentInput = append(p.currentInput, char)
			p.position++
		}
	}
}
