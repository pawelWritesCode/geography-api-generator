package expand

import (
	"generator/backend-go/templates"
	"generator/backend-go/tools/generator"
	"log"
	"sync"
)

type Worker struct{}

//New returns new Worker struct
func New() Worker {
	return Worker{}
}

var wg sync.WaitGroup

//ExpandRandom expands project by one random entity
func (w Worker) ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty) error {
	randomVariables, err := generator.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
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
