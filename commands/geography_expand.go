// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/tools/expand"
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"github.com/urfave/cli/v2"
)

//ErrInvalidDirectoryStructure occurs when there are missing some folders. Probably user is not in geography root
var ErrInvalidDirectoryStructure = errors.New("invalid directory structure")

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error {
	err := checkDirectoryStructure()

	if errors.Is(err, ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	}

	worker := expand.New()
	eGen := generator.NewEntityGenerator()
	pGen := generator.NewPropertyGenerator()
	err = worker.ExpandRandom(eGen, pGen)

	if err != nil {
		return err
	}
	return nil
}

//checkDirectoryStructure checks if user is in geography root folder
func checkDirectoryStructure() error {
	dirs := [7]resource.Resource{
		resource.New(geography.EntityDir, ""),
		resource.New(geography.ControllerDir, ""),
		resource.New(geography.RepositoryDir, ""),
		resource.New(geography.ResourcesDir, ""),
		resource.New(geography.RestApiDir, ""),
		resource.New(geography.BehatDir, ""),
		resource.New(geography.DocumentationDir, ""),
	}

	for _, dir := range dirs {
		if !dir.Exist() {
			return ErrInvalidDirectoryStructure
		}
	}

	return nil
}
