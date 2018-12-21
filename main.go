package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "jk-write"
	app.Usage = "Write gh-page(made by Jekyll) post"
	app.Version = "0.1"
	app.Author = "shamisonn"
	app.Email = "sh4m1s0nn@gmail.com"
	app.Commands = Commands
	return app
}
