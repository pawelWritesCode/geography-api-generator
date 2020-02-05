package commands

import (
	"generator/backend-go/generators"
	"generator/backend-go/templates"
	"github.com/urfave/cli/v2"
)

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error  {
	//TODO: check if user is in right folder.

	randomTemplateVariables, err := generators.RandomTemplateVariables(10)

	if err != nil {
		return err
	}

	allTemplates  := []templates.Template{
		templates.NewEntity(randomTemplateVariables),
	}

	for _, tpl := range allTemplates {
		err := tpl.RenderAndEmplace()

		//TODO: Remove already created files, before returning error.
		if err != nil {
			return err
		}
	}

	return nil
}