package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/generators"
	"generator/backend-go/templates"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sync"
)

//ErrInvalidDirectoryStructure occurs when there are missing some folders. Probably user is not in geography root
var ErrInvalidDirectoryStructure = errors.New("invalid directory structure")

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error {
	err := checkDirectoryStructure()

	if errors.Is(err, ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	}

	var entity generators.Entity
	var property generators.Property
	randomVariables, err := generators.RandomTemplateVariables(entity, property, 10)

	if errors.Is(err, generators.ErrExpand) {
		return fmt.Errorf("%v, project cannot expand anymore", err)
	}

	allTemplates := []templates.Template{
		templates.NewEntity(randomVariables),
		templates.NewControllerDelete(randomVariables),
		templates.NewControllerGet(randomVariables),
		templates.NewControllerGetList(randomVariables),
		templates.NewControllerPost(randomVariables),
		templates.NewControllerPut(randomVariables),
		templates.NewResource(randomVariables),
		templates.NewRepository(randomVariables),
		templates.NewRestApiDelete(randomVariables),
		templates.NewRestApiGetList(randomVariables),
		templates.NewRestApiPost(randomVariables),
		templates.NewRestApiPut(randomVariables),
		templates.NewBehatCreate(randomVariables),
		templates.NewBehatGetId(randomVariables),
		templates.NewBehatDelete(randomVariables),
		templates.NewBehatGetList(randomVariables),
		templates.NewBehatPut(randomVariables),
		templates.NewDocumentationRequest(randomVariables),
		templates.NewDocumentationResponseSingle(randomVariables),
		templates.NewDocumentationResponseArray(randomVariables),
	}
	var wg sync.WaitGroup

	for _, tpl := range allTemplates {
		wg.Add(1)

		go func(tpl templates.Template) {
			defer wg.Done()
			err := tpl.RenderAndEmplace()

			if err != nil {
				log.Fatal(err)
			}
		}(tpl)

	}

	wg.Wait()
	return nil
}

//checkDirectoryStructure checks if user is in geography root folder
func checkDirectoryStructure() error {

	_, err := os.Stat(templates.EntityDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.ControllerDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.RepositoryDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.ResourcesDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.RestApiDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.BehatDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	_, err = os.Stat(templates.DocumentationDirectory)
	if os.IsNotExist(err) {
		return ErrInvalidDirectoryStructure
	}

	return nil
}
