// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/decay"
	"generator/backend-go/tools/decay/picker"
	"generator/backend-go/tools/geography"
	"github.com/urfave/cli/v2"
)

//GeographyExpand command shrinks geography project by one random entity and related to it files like controllers etc..
func GeographyShrink(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	} else if err != nil {
		return err
	}

	randomEntityPicker := picker.New()

	worker := decay.New()
	err = worker.ShrinkRandom(randomEntityPicker)

	if errors.Is(err, picker.ErrNoAvailableEntities) {
		return fmt.Errorf("there are no entities left for shrinking project")
	} else if err != nil {
		return err
	}

	return nil
}
