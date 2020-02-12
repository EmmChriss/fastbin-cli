package main

import (
	"log"
	"os"

	"github.com/jozsefsallai/fastbin-cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fastbin"
	app.Usage = "command line client for fastbin"
	app.Version = "1.0.0"
	app.Action = commands.CreateSnippet

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
