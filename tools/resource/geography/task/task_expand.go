package task

import (
	"generator/backend-go/tools"
	"generator/backend-go/tools/resource/geography/templates"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
)

//ExpandRandom expands project by one random entity
func (t Task) ExpandRandom(e tools.Employee, eGen generator.RandomEntity, pGen generator.RandomProperty) error {
	randomVariables, err := generator.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
	}

	return t.ExpandSpecific(e, randomVariables)
}

//ExpandSpecific expands project by one entity
func (t Task) ExpandSpecific(e tools.Employee, randomVariables templateUtils.TemplateVariables) error {
	e.RegisterJob(templates.NewEntity(randomVariables))
	e.RegisterJob(templates.NewControllerGet(randomVariables))
	e.RegisterJob(templates.NewControllerDelete(randomVariables))
	e.RegisterJob(templates.NewControllerGetList(randomVariables))
	e.RegisterJob(templates.NewControllerPost(randomVariables))
	e.RegisterJob(templates.NewControllerPut(randomVariables))
	e.RegisterJob(templates.NewResource(randomVariables))
	e.RegisterJob(templates.NewRepository(randomVariables))
	e.RegisterJob(templates.NewRestApiDelete(randomVariables))
	e.RegisterJob(templates.NewRestApiGetList(randomVariables))
	e.RegisterJob(templates.NewRestApiPost(randomVariables))
	e.RegisterJob(templates.NewRestApiPut(randomVariables))
	e.RegisterJob(templates.NewBehatCreate(randomVariables))
	e.RegisterJob(templates.NewBehatGetId(randomVariables))
	e.RegisterJob(templates.NewBehatDelete(randomVariables))
	e.RegisterJob(templates.NewBehatGetList(randomVariables))
	e.RegisterJob(templates.NewBehatPut(randomVariables))
	e.RegisterJob(templates.NewDocumentationRequest(randomVariables))
	e.RegisterJob(templates.NewDocumentationResponseSingle(randomVariables))
	e.RegisterJob(templates.NewDocumentationResponseArray(randomVariables))

	return e.DoAll()
}
