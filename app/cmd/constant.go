package cmd

import "os"

type CommandHandler func(args []string, outFile *os.File, errFile *os.File)

func init() {
	// Initialize the list of all commands
	AllCommands = make([]string, 0, len(HandlerMap))
	for cmd := range HandlerMap {
		AllCommands = append(AllCommands, cmd)
	}
}

var AllCommands = []string{}

var HandlerMap = map[string]CommandHandler{
	"exit": HandleExit,
	"echo": HandleEcho,
	"type": HandleType,
	"pwd":  HandlePWD,
	"cd":   HandleCD,
}
