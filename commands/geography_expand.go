// Package commands contains generator commands
package commands

import (
	"errors"
	"fmt"
	"generator/backend-go/templates"
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"github.com/urfave/cli/v2"
	"log"
	"sync"
)

var wg sync.WaitGroup

//ErrInvalidDirectoryStructure occurs when there are missing some folders. Probably user is not in geography root
var ErrInvalidDirectoryStructure = errors.New("invalid directory structure")

//GeographyExpand command updates geography project by new entity and related to it files like controllers etc..
func GeographyExpand(c *cli.Context) error {
	err := checkDirectoryStructure()

	if errors.Is(err, ErrInvalidDirectoryStructure) {
		return fmt.Errorf("%v, change directory to geography root directory", err)
	}

	var entity generator.Entity
	var property generator.Property
	randomVariables, err := generator.RandomTemplateVariables(entity, property, 10)

	if errors.Is(err, generator.ErrExpand) {
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

	for _, tpl := range allTemplates {
		wg.Add(1)
		go renderAndWrite(tpl)
	}

	wg.Wait()
	return nil
}

//renderAndWrite renders template and emplace it.
func renderAndWrite(tpl templates.Template) {
	defer wg.Done()
	err := tpl.RenderAndWrite()

	if err != nil {
		log.Fatal(err)
	}
}

//checkDirectoryStructure checks if user is in geography root folder
func checkDirectoryStructure() error {
	entityDirectory := resource.New(geography.EntityDir, "")
	if !entityDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	controllerDirectory := resource.New(geography.ControllerDir, "")
	if !controllerDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	repositoryDirectory := resource.New(geography.RepositoryDir, "")
	if !repositoryDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	resourceDirectory := resource.New(geography.ResourcesDir, "")
	if !resourceDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	restApiDirectory := resource.New(geography.RestApiDir, "")
	if !restApiDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	behatDirectory := resource.New(geography.BehatDir, "")
	if !behatDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	documentationDirectory := resource.New(geography.DocumentationDir, "")
	if !documentationDirectory.Exist() {
		return ErrInvalidDirectoryStructure
	}

	return nil
}
