//Package expand implement methods for expanding project.
//
//To expand project, instantiate new worker using
//	NewWorkerExpand()
//to expand by one random entity use method
//	ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty)
package worker

import (
	"context"
	templates2 "generator/backend-go/tools/resource/geography/templates"
	templateUtils2 "generator/backend-go/tools/resource/geography/templates/templateUtils"
	generator2 "generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
)

//WorkerExpand is responsible for expanding project
type WorkerExpand struct{}

//NewWorkerExpand returns new WorkerExpand struct
func NewWorkerExpand() WorkerExpand {
	return WorkerExpand{}
}

//ExpandRandom expands project by one random entity
func (w WorkerExpand) ExpandRandom(eGen generator2.RandomEntity, pGen generator2.RandomProperty) error {
	randomVariables, err := generator2.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
	}

	return w.ExpandSpecific(randomVariables)
}

//ExpandSpecific expands project by one entity
func (w WorkerExpand) ExpandSpecific(randomVariables templateUtils2.TemplateVariables) error {
	allTemplates := []templates2.Template{
		templates2.NewEntity(randomVariables),
		templates2.NewControllerDelete(randomVariables),
		templates2.NewControllerGet(randomVariables),
		templates2.NewControllerGetList(randomVariables),
		templates2.NewControllerPost(randomVariables),
		templates2.NewControllerPut(randomVariables),
		templates2.NewResource(randomVariables),
		templates2.NewRepository(randomVariables),
		templates2.NewRestApiDelete(randomVariables),
		templates2.NewRestApiGetList(randomVariables),
		templates2.NewRestApiPost(randomVariables),
		templates2.NewRestApiPut(randomVariables),
		templates2.NewBehatCreate(randomVariables),
		templates2.NewBehatGetId(randomVariables),
		templates2.NewBehatDelete(randomVariables),
		templates2.NewBehatGetList(randomVariables),
		templates2.NewBehatPut(randomVariables),
		templates2.NewDocumentationRequest(randomVariables),
		templates2.NewDocumentationResponseSingle(randomVariables),
		templates2.NewDocumentationResponseArray(randomVariables),
	}

	ch1 := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
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
func renderAndWrite(ctx context.Context, tpl templates2.Template, ch1 chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		ch1 <- tpl.RenderAndWrite()
	}
}
