//Package expand implement methods for expanding project.
//
//To expand project, instantiate new worker using
//	New()
//to expand by one random entity use method
//	ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty)
package expand

import (
	"context"
	"generator/backend-go/templates"
	"generator/backend-go/tools/generator"
)

//Worker is responsible for expanding project
type Worker struct{}

//New returns new Worker struct
func New() Worker {
	return Worker{}
}

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

	ch1 := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, tpl := range allTemplates {
		go renderAndWrite(ctx, tpl, ch1)
	}

	for i := 0; i < len(allTemplates); i++ {
		err = <-ch1

		if err != nil {
			cancel()
			return err
		}
	}

	return nil
}

//renderAndWrite renders template and emplace it.
func renderAndWrite(ctx context.Context, tpl templates.Template, ch1 chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		ch1 <- tpl.RenderAndWrite()
	}
}
