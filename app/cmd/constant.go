package cmd

import "os"

// CommandHandler is a function type that handles a command.
// It returns true if the executed command printed something to the output or error file, and false otherwise.
// The handler should not print any NewLine or CRLF characters to the output or error file.
type CommandHandler func(args []string, outFile *os.File, errFile *os.File) bool

func init() {
	// Initialize the list of all commands
	AllCommands = make([]string, 0, len(HandlerMap))
	for cmd := range HandlerMap {
		AllCommands = append(AllCommands, cmd)
	}

	// Manually register the history command
	AllCommands = append(AllCommands, "history")
}

var AllCommands = []string{}

var HandlerMap = map[string]CommandHandler{
	"exit": HandleExit,
	"echo": HandleEcho,
	"type": HandleType,
	"pwd":  HandlePWD,
	"cd":   HandleCD,
}
