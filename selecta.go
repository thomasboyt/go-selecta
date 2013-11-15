package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "selecta"
	app.Usage = "fuzzy find whatever you want"
	app.Action = func(c *cli.Context) {
		println("Hello!")
	}

	app.Run(os.Args)
}
