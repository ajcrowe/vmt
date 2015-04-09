package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// create and configure app
	app := cli.NewApp()
	app.Name = "vmt"
	app.Version = "0.1.0"
	app.Author = "Alex Crowe"
	app.Usage = "Generate and modify Vagrant box json manifests"

	// load config
	loadConfig()

	// load commands
	app.Commands = Commmands
	app.Run(os.Args)
}
