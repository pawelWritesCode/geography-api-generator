package main

import (
	"generator/backend-go/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "geo-generator",
		Usage: "CLI for maintaining geography project.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "display verbose output",
				Aliases: []string{"v"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "geography:expand",
				Usage:  " ⇗ expand project by one entity",
				Action: commands.GeographyExpand,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}},
				},
			},
			{
				Name:   "geography:shrink",
				Usage:  " ⇘ shrink project by one entity",
				Action: commands.GeographyShrink,
			},
			{
				Name:   "geography:rename",
				Usage:  " ⇔ rename one entity and all its related files/folders",
				Action: commands.GeographyRename,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
