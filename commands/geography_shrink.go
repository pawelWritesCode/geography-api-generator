// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	worker2 "generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
	"log"
)

//GeographyExpand command shrinks geography project by one random entity and related to it files like controllers etc..
func GeographyShrink(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("⛔ %v, change directory to geography root directory", err)
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Valid directory structure ✓ shrinking project by one entity")
	}
	randomEntityPicker := picker.New()

	worker := worker2.NewWorkerDecay()
	err = worker.ShrinkRandom(randomEntityPicker)

	if errors.Is(err, picker.ErrNoAvailableEntities) {
		return fmt.Errorf("⛔ there are no entities left for shrinking project")
	} else if err != nil {
		return fmt.Errorf("⛔ %v", err)
	}

	if c.Bool("verbose") {
		log.Println("Shrinking has succeeded")
	}

	return nil
}
