// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	worker2 "generator/backend-go/tools/resource/geography/worker"
	"github.com/urfave/cli/v2"
)

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error {
	err := geography.CheckDirStructure()

	if errors.Is(err, geography.ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	} else if err != nil {
		return err
	}

	worker := worker2.NewWorkerExpand()
	eGen := generator.NewEntityGenerator()
	pGen := generator.NewPropertyGenerator()
	err = worker.ExpandRandom(eGen, pGen)

	if errors.Is(err, generator.ErrExpand) {
		return fmt.Errorf("%v, project cannot expand anymore", err)
	} else if err != nil {
		return err
	}

	return nil
}
