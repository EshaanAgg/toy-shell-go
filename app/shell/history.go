package shell

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Appends the history loaded from the specified file to the shell's history.
// Also updates the last saved history index to the end of the loaded history.
// If addCurrentCmd is true, it adds the current command to the history before loading.
func (s *Shell) loadHistory(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read history file: %w", err)
	}

	// Split the content by newlines to get the history commands
	// and remove the last empty line if it exists.
	s.history = append(s.history, strings.Split(string(content), "\n")...)
	if len(s.history) > 0 && s.history[len(s.history)-1] == "" {
		s.history = s.history[:len(s.history)-1]
	}
	s.lastSavedHistoryIdx = len(s.history) - 1

	return nil
}

// Saves the history to the specified file.
// It appends new commands to the file, starting from the last saved index.
// If the file does not exist, it creates a new one.
func (s *Shell) saveHistory(filePath string) error {
	content := strings.Join(s.history[s.lastSavedHistoryIdx+1:], "\n")

	// Open the file in Append mode, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open history file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to history file: %w", err)
	}

	s.lastSavedHistoryIdx = len(s.history) - 1
	return nil
}

// Handler for the history command
func (c *command) handleHistory(s *Shell) {
	// Handle -w, -r and -a flags for history command
	if len(c.args) == 3 {
		switch c.args[1] {
		case "-w", "-a":
			if err := s.saveHistory(c.args[2]); err != nil {
				fmt.Fprintf(c.errFile, "Error saving history: %v\r\n", err)
			}
			return
		case "-r":
			if err := s.loadHistory(c.args[2]); err != nil {
				fmt.Fprintf(c.errFile, "Error loading history: %v\r\n", err)
			}
			return
		default:
			fmt.Fprintf(c.errFile, "Unknown flag: %s\r\n", c.args[1])
			return
		}
	}

	// Handle printing the history
	startIdx := 0
	if len(c.args) > 1 {
		cmdCnt, err := strconv.Atoi(c.args[1])
		if err != nil || cmdCnt < 0 {
			fmt.Fprintf(c.errFile, "Invalid argument: %s\r\n", c.args[1])
			return
		}

		startIdx = len(s.history) - cmdCnt
	}

	for i, cmd := range s.history[startIdx:] {
		fmt.Fprintf(c.outFile, "\t%d %s\r\n", i+1, cmd)
	}
}

// Clear's the current line in the shell
func (s *Shell) clearLine() {
	fmt.Printf("\r%s", MOVE_CURSOR_LEFT)
	fmt.Print(CLEAR_LINE)
	fmt.Print(MOVE_CURSOR_LEFT)
}

// Set's the current command to from the history
func (s *Shell) putFromHistory() {
	if s.curHistoryIdx == -1 || s.curHistoryIdx >= len(s.history) {
		return
	}

	s.clearLine()
	s.input = []byte(s.history[s.curHistoryIdx])
	s.cursorPosition = len(s.input)
	fmt.Printf("$ %s", s.input)
}

func (s *Shell) handleUpArrowPress() {
	if s.curHistoryIdx == -1 {
		s.curHistoryIdx = len(s.history) - 1
	} else if s.curHistoryIdx > 0 {
		s.curHistoryIdx--
	}
	s.putFromHistory()
}

func (s *Shell) handleDownArrowPress() {
	if s.curHistoryIdx == -1 {
		return
	}

	if s.curHistoryIdx < len(s.history)-1 {
		s.curHistoryIdx++
	} else {
		s.curHistoryIdx = -1
	}

	s.putFromHistory()
}
