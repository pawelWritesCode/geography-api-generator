// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/decay"
	"generator/backend-go/tools/decay/picker"
	"github.com/urfave/cli/v2"
)

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyShrink(c *cli.Context) error {
	err := checkDirectoryStructure()

	if errors.Is(err, ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	}

	geographyPicker := picker.Picker{}
	entity, err := geographyPicker.RandomEntity()
	fmt.Println("entity " + entity)

	if errors.Is(err, picker.ErrNoAvailableEntities) {
		return fmt.Errorf("there are no entities left for shrinking project")
	} else if err != nil {
		return err
	}

	worker := decay.Worker{}
	err = worker.Shrink(entity)

	if err != nil {
		return err
	}

	return nil
}
