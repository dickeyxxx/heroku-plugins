package plugins

import "github.com/dickeyxxx/heroku-plugins/cli"

var cmdList = &cli.Command{
	ShortHelp: "Lists the installed plugins",
	Help: `Lists installed plugins

  Example:
  $ heroku plugins`,

	Run: func(ctx *cli.Context) {
		packages, err := node.Packages()
		must(err)
		for _, pkg := range packages {
			cli.Println(pkg.Name, pkg.Version)
		}
	},
}
