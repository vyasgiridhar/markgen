package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/vyasgiridhar/markgen"
)

func main() {
	app := cli.NewApp()
	app.Name = "markgen"
	app.Version = markgen.Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 6060,
			Usage: "Port to listen.",
		},
	}
	app.Action = func(c *cli.Context) {
		args := c.Args()

		markgen := markgen.NewMarkgen(c.Int("port"))
		markgen.Run(args...)
	}
	app.Run(os.Args)

}
