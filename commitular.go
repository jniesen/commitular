package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "commitular"
	app.Usage = "Make a git commit using the Angular commit message format"

	app.Run(os.Args)
}
