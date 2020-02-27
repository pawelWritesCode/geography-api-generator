package worker

import (
	"context"
	"generator/backend-go/tools/resource/geography/templates"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
)

//WorkerExpand is responsible for expanding project
type WorkerExpand struct{}

//NewWorkerExpand returns new WorkerExpand struct
func NewWorkerExpand() WorkerExpand {
	return WorkerExpand{}
}

//ExpandRandom expands project by one random entity
func (w WorkerExpand) ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty) error {
	randomVariables, err := generator.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
	}

	return w.ExpandSpecific(randomVariables)
}

//ExpandSpecific expands project by one entity
func (w WorkerExpand) ExpandSpecific(randomVariables templateUtils.TemplateVariables) error {
	allTemplates := []templates.Template{
		templates.NewEntity(randomVariables),
		templates.NewControllerGet(randomVariables),
		templates.NewControllerDelete(randomVariables),
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
func renderAndWrite(ctx context.Context, tpl templates.Template, ch1 chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		ch1 <- tpl.RenderAndWrite()
	}
}
