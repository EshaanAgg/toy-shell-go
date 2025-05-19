package utils

import (
	"bytes"
	"fmt"
)

type parser struct {
	idx   int
	input []byte

	tokens       []string
	currentToken string
	escapeNext   bool

	// If nil, then the parser is not in a quoted string.
	// If not nil, then the parser is in a quoted string and this byte
	// is the quote character.
	quoted *byte
}

func (p *parser) next() byte {
	if p.idx >= len(p.input) {
		return 0
	}
	b := p.input[p.idx]
	p.idx++
	return b
}

func (p *parser) expectDelimiter() error {
	b := p.next()
	if b == 0 {
		return nil
	}
	if b != ' ' && b != '\t' && b != '\n' && b != '\r' {
		return fmt.Errorf("expected a delimeter to seperate tokens at %d, got %q", p.idx-1, b)
	}
	return nil
}

// addToken adds the current token to the list of tokens, only if
// the current token is not empty. It also resets the current token.
func (p *parser) addToken() {
	if p.currentToken != "" {
		p.tokens = append(p.tokens, p.currentToken)
		p.currentToken = ""
	}
}

func (p *parser) inSingleQuotes() bool {
	return p.quoted != nil && *p.quoted == '\''
}

func (p *parser) inDoubleQuotes() bool {
	return p.quoted != nil && *p.quoted == '"'
}

func (p *parser) parse() []string {
	for {
		b := p.next()

		if b == 0 {
			p.addToken()
			return p.tokens
		}

		if p.inSingleQuotes() {
			p.handleSingleQuote(b)
			continue
		}

		if p.inDoubleQuotes() {
			p.handleDoubleQuote(b)
			continue
		}

		// Now we are not in a quoted string, thus we only need to
		// check for escape characters, and start of a quoted string.

		if p.escapeNext {
			p.currentToken += string(b)
			p.escapeNext = false
			continue
		}

		switch b {
		case '\\':
			p.escapeNext = true
		case '"':
			p.quoted = &b
		case '\'':
			p.quoted = &b

		// Only space and tabs are considered as delimiters for tokens
		case ' ', '\t':
			p.addToken()
		default:
			p.currentToken += string(b)
		}
	}
}

// In single quote, all characters are part of the token.
// Only a single quote can end the token.
// New token should NOT BE created when the single quote is closed.
func (p *parser) handleSingleQuote(b byte) {
	if b == '\'' {
		p.quoted = nil
	} else {
		p.currentToken += string(b)
	}
}

// In double quotes, there is limited escaping.
// The blackslash (\) can be used to escape (' " $ \n) bytes.
// New token should NOT BE created when the double quote is closed.
func (p *parser) handleDoubleQuote(b byte) {
	if p.escapeNext {
		switch b {
		case '\\', '"', '$', '\n':
			p.currentToken += string(b)
		default:
			// Preserve the blackslash that was earlier used to escape
			p.currentToken += "\\" + string(b)
		}
		p.escapeNext = false
		return
	}

	switch b {
	case '\\':
		p.escapeNext = true

	case '"':
		p.quoted = nil

	default:
		p.currentToken += string(b)
	}
}

// GetTokens parses the input byte array and returns a slice of tokens.
// It carefully handles quoted strings and escape characters, and returns the
// final tokens that can be used to execute a command.
func GetTokens(input []byte) ([]string, error) {
	p := &parser{
		input: bytes.TrimSpace(input),
	}

	toks := p.parse()
	if p.inSingleQuotes() || p.inDoubleQuotes() {
		return nil, fmt.Errorf("unmatched quotes in input")
	}
	if p.escapeNext {
		return nil, fmt.Errorf("unmatched escape character in input")
	}
	return toks, nil
}
