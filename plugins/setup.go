package plugins

import "github.com/dickeyxxx/heroku-plugins/cli"

func Setup() {
	if node.IsSetup() {
		return
	}
	cli.Err("setting up plugins... ")
	must(node.Setup())
	cli.Errln("done")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
