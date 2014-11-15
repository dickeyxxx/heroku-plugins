package main

var commands = &Topic{
	Name:      "commands",
	ShortHelp: "list all commands",
	Commands:  []*Command{commandsRun},
}

var commandsRun = &Command{
	ShortHelp: "list all commands",
	Run: func(ctx *Context) {
		for _, topic := range PluginTopics() {
			for _, command := range topic.Commands {
				Printf("%s:%s\n", topic.Name, command.Name)
			}
		}
	},
}
