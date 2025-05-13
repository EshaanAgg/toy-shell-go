package shell

import "fmt"

func (s *Shell) resetState() {
	s.input = nil
	s.cursorPosition = 0
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
			s.ExecuteCommand(s.input)

			s.resetState()
			s.putPrompt()

		default:
			s.input = append(s.input, ch)
			s.cursorPosition++
			fmt.Printf("%c", ch)
		}
	}
}
