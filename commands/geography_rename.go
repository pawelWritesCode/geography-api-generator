// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	"generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
)

//GeographyRename command is responsible for renaming one random entity
func GeographyRename(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	} else if err != nil {
		return err
	}

	randomPicker := picker.New()
	eGen := generator.NewEntityGenerator()
	workerRename := worker.NewWorkerRename()
	err = workerRename.RenameRandom(eGen, randomPicker)

	if errors.Is(err, picker.ErrNoAvailableEntities) {
		return fmt.Errorf("there are no entities left for renaming")
	} else if err != nil {
		return err
	}

	return nil
}
