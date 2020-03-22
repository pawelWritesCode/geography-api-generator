// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	t "generator/backend-go/tools/resource/geography/task"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	"generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
	"log"
)

//GeographyRename command is responsible for renaming one random entity
func GeographyRename(c *cli.Context) error {
	allGeoDirs := geography.AllGeographyDirectories()
	sAllGeoDirs := allGeoDirs[:]
	err := resource.CheckDirStructure(sAllGeoDirs)

	if errors.Is(err, resource.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("⛔ %v, change directory to geography root directory", err)
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Valid directory structure ✓ renaming one entity")
	}

	task := t.New()
	err = task.RenameRandomToRandom(worker.New(), generator.NewEntityGenerator(), picker.New())

	if errors.Is(err, picker.ErrNoAvailableEntities) {
		return fmt.Errorf("⛔ there are no entities left for renaming")
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Renaming has succeeded")
	}

	return nil
}
