// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	t "generator/backend-go/tools/resource/geography/task"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	"github.com/urfave/cli/v2"
	"log"
)

//GeographyRename command is responsible for renaming one random entity
func GeographyRename(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("⛔ %v, change directory to geography root directory", err)
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Valid directory structure ✓ renaming one entity")
	}

	randomPicker := picker.New()
	eGen := generator.NewEntityGenerator()
	task := t.New()
	err = task.RenameRandom(eGen, randomPicker)

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
