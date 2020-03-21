package task

import (
	"generator/backend-go/tools/resource/geography/templates"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/worker"
)

//ExpandRandom expands project by one random entity
func (t Task) ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty) error {
	randomVariables, err := generator.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
	}

	return t.ExpandSpecific(randomVariables)
}

//ExpandSpecific expands project by one entity
func (t Task) ExpandSpecific(randomVariables templateUtils.TemplateVariables) error {
	w := worker.NewWorker()

	w.RegisterJob(templates.NewEntity(randomVariables))
	w.RegisterJob(templates.NewControllerGet(randomVariables))
	w.RegisterJob(templates.NewControllerDelete(randomVariables))
	w.RegisterJob(templates.NewControllerGetList(randomVariables))
	w.RegisterJob(templates.NewControllerPost(randomVariables))
	w.RegisterJob(templates.NewControllerPut(randomVariables))
	w.RegisterJob(templates.NewResource(randomVariables))
	w.RegisterJob(templates.NewRepository(randomVariables))
	w.RegisterJob(templates.NewRestApiDelete(randomVariables))
	w.RegisterJob(templates.NewRestApiGetList(randomVariables))
	w.RegisterJob(templates.NewRestApiPost(randomVariables))
	w.RegisterJob(templates.NewRestApiPut(randomVariables))
	w.RegisterJob(templates.NewBehatCreate(randomVariables))
	w.RegisterJob(templates.NewBehatGetId(randomVariables))
	w.RegisterJob(templates.NewBehatDelete(randomVariables))
	w.RegisterJob(templates.NewBehatGetList(randomVariables))
	w.RegisterJob(templates.NewBehatPut(randomVariables))
	w.RegisterJob(templates.NewDocumentationRequest(randomVariables))
	w.RegisterJob(templates.NewDocumentationResponseSingle(randomVariables))
	w.RegisterJob(templates.NewDocumentationResponseArray(randomVariables))

	return w.DoAll()
}
