package shell

import (
	"fmt"
	"os"
)

func (s *Shell) resetState() {
	s.input = nil
	s.cursorPosition = 0
}

func (s *Shell) putChar(ch byte) {
	s.input = append(s.input, ch)
	s.cursorPosition++
	fmt.Printf("%c", ch)
}

func (s *Shell) Start() []byte {
	s.putPrompt()

	for {
		ch := s.readByte()

		switch ch {
		case KEY_CTRL_C:
			fmt.Printf("\r\nExiting shell...\r\n")
			s.Kill()

		case KEY_ENTER, KEY_NEWLINE:
			fmt.Printf("\r\n")
			s.history = append(s.history, string(s.input))
			s.curHistoryIdx = -1
			s.ExecuteCommand(s.input)

			s.resetState()
			s.putPrompt()

		case KEY_BACKSPACE:
			s.handleBackspace()

		case KEY_ESC:
			// Read 2 more bytes
			var seq [2]byte
			os.Stdin.Read(seq[:])

			if seq[0] == '[' {
				switch seq[1] {
				case 'A':
					s.handleUpArrowPress()
				case 'B':
					s.hanldeDownArrowPress()
				}
			}

		case KEY_TAB:
			s.handleTabClick()

		default:
			s.putChar(ch)
		}

		// Reset the state if the user has typed
		// something other than TAB after multiple matches
		if ch != KEY_TAB && s.hadMultipleMatched {
			s.hadMultipleMatched = false
		}
	}
}

func (s *Shell) handleBackspace() {
	if s.cursorPosition == 0 {
		return
	}

	// Update the internal state of the shell
	s.cursorPosition--
	s.input = append(s.input[:s.cursorPosition], s.input[s.cursorPosition+1:]...)

	// Redraw the rest of the input from the cursor position
	// Add a space to overwrite the trailing character
	toPrintAgain := string(s.input[s.cursorPosition:]) + " "

	fmt.Print(MOVE_CURSOR_LEFT)
	fmt.Print(toPrintAgain)
	// Move cursor back to its proper position
	for _ = range len(toPrintAgain) {
		fmt.Print(MOVE_CURSOR_LEFT)
	}
}
