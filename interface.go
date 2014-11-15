package main

import (
	"encoding/json"
	"fmt"
)

func runFn(module, topic, command string) func(ctx *Context) {
	return func(ctx *Context) {
		ctxJson, err := json.Marshal(ctx)
		must(err)
		script := fmt.Sprintf(`
		require('%s')
		.topics.filter(function (topic) {
			return topic.name == '%s'
		})[0]
		.commands.filter(function (command) {
			return command.name == '%s'
		})[0]
		.run(%s)`, module, topic, command, ctxJson)

		cmd := node.RunScript(script)
		cmd.Stdout = Stdout
		cmd.Stderr = Stderr
		must(cmd.Run())
	}
}

func getPackageTopics(name string) []*Topic {
	script := `console.log(JSON.stringify(require('` + name + `')))`
	cmd := node.RunScript(script)
	cmd.Stderr = Stderr
	output, err := cmd.StdoutPipe()
	must(err)
	must(cmd.Start())
	var response map[string][]*Topic
	must(json.NewDecoder(output).Decode(&response))
	must(cmd.Wait())
	topics := response["topics"]
	for _, topic := range topics {
		for _, command := range topic.Commands {
			command.Run = runFn(name, topic.Name, command.Name)
		}
	}
	return topics
}

func PluginTopics() (topics []*Topic) {
	packages, err := node.Packages()
	must(err)
	for _, pkg := range packages {
		topics = append(topics, getPackageTopics(pkg.Name)...)
	}
	return topics
}
