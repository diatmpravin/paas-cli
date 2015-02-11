package main

import (
	"esp-cli/esp/configuration"
	term "esp-cli/esp/terminal"
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
				config := configuration.Default()
				fmt.Println("Target Information (where will apps be pushed):")
				fmt.Printf("CF instance %s (API Version %s)\n",
					term.Colorize(config.Target, term.Yellow, true),
					term.Colorize(config.ApiVersion, term.Yellow, true))
				fmt.Println("user: N/A")
				fmt.Println("target app space: N/A (org: N/A)")
			},
		},
	}

	app.Run(os.Args)
}
