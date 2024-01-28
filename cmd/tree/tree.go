package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go-demos/tree"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "List contents of directories in a tree-like format"
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:    "level",
			Aliases: []string{"l"},
			Usage:   "Max depth to print",
		},
	}
	app.Action = func(c *cli.Context) error {
		dir := "./"
		if c.NArg() > 0 {
			dir = c.Args().Get(0)
		}
		//level := c.Int("level")

		return tree.ScanWalkDir(dir, 0)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
