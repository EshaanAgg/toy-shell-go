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

// cleanup closes the provided files and ignores any errors.
// It filters out stdout and stderr to avoid closing them.
// This is useful for cleaning up file descriptors after command execution to avoid resource leaks.
func cleanup(files ...*os.File) {
	for _, file := range files {
		if file != os.Stdout && file != os.Stderr {
			file.Close()
		}
	}
}
