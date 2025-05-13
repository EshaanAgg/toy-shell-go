package cmd

type CommandHandler func(args []string)

var AllCommands = []string{
	"exit",
	"echo",
	"type",
	"pwd",
	"cd",
}

var HandlerMap = map[string]CommandHandler{
	"exit": HandleExit,
	"echo": HandleEcho,
	"type": HandleType,
	"pwd":  HandlePWD,
	"cd":   HandleCD,
}
