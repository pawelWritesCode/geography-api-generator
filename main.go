package main

import (
	"generator/backend-go/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "geography:expand",
				Usage:  "expand project by one entity",
				Action: commands.GeographyExpand,
			},
			{
				Name:   "geography:shrink",
				Usage:  "shrink project by one entity",
				Action: commands.GeographyShrink,
			},
			{
				Name:   "geography:rename",
				Usage:  "rename one entity and all its related files/folders",
				Action: commands.GeographyRename,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
