// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	t "generator/backend-go/tools/resource/geography/task"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
	"log"
)

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("⛔ %v, change directory to geography root directory", err)
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Valid directory structure ✓ expanding project by one entity")
	}

	task := t.New()
	err = task.ExpandRandom(worker.New(), generator.NewEntityGenerator(), generator.NewPropertyGenerator())

	if errors.Is(err, generator.ErrExpand) {
		return fmt.Errorf("⛔ %v, project cannot expand anymore", err)
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Expanding has succeeded")
	}

	return nil
}
