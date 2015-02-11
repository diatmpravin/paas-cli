package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "esp"
	app.Usage = "A command line tool to interact with ESP"
	app.Version = "1.0.0.alpha"
	app.Commands = []cli.Command{
		{
			Name:      "target",
			ShortName: "t",
			Usage:     "Set or View the target",
			Action: func(c *cli.Context) {
				fmt.Println("Target Information (where will apps be pushed):")

			},
		},
	}

	app.Run(os.Args)
}
