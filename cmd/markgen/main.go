package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/vyasgiridhar/markgen"
)

func main() {
	app := cli.NewApp()
	app.Name = "markgen"
	app.Version = markgen.Version
	app.Usage = `markgen is a Markdown previewer written in Go.
   For information, please visit https://github.com/vyasgiridhar/markgen`
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 5000,
			Usage: "Port to listen in",
		},
	}
	app.Action = func(c *cli.Context) {
		args := c.Args()

		markgen := markgen.NewMarkGen(c.Int("port"))

		if c.Bool("basic") {
			markgen.UseBasic()
		}

		markgen.Run(args...)
	}

	app.Run(os.Args)
}
