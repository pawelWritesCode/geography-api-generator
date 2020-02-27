// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	"generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
)

//GeographyExpand command shrinks geography project by one random entity and related to it files like controllers etc..
func GeographyRename(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	} else if err != nil {
		return err
	}

	randomPicker := picker.New()
	rnd, err := randomPicker.RandomEntityAndProperty()

	if err != nil {
		return err
	}
	//Remove this entity
	decayWorker := worker.NewWorkerDecay()
	err = decayWorker.ShrinkSpecific(rnd.Entity)

	if err != nil {
		return err
	}

	//Add new one with same property
	entityGenerator := generator.NewEntityGenerator()
	specificVariables := templateUtils.NewTemplateVariables(entityGenerator.Random(), rnd.Property)

	expandWorker := worker.NewWorkerExpand()
	err = expandWorker.ExpandSpecific(specificVariables)

	if err != nil {
		return err
	}

	return nil
}
