package cmd

type CommandHandler func(args []string)

var AllCommands = []string{
	"exit",
	"echo",
	"type",
}

var HandlerMap = map[string]CommandHandler{
	"exit": HandleExit,
	"echo": HandleEcho,
	"type": HandleType,
}
