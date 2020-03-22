package task

import (
	"generator/backend-go/tools"
	"generator/backend-go/tools/resource/geography/templates"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
)

//ExpandRandom expands project by one random entity
func (t Task) ExpandRandom(e tools.Employee, eGen generator.RandomEntity, pGen generator.RandomProperty) error {
	randomVariables, err := generator.RandomTemplateVariables(eGen, pGen, 10)

	if err != nil {
		return err
	}

	return t.ExpandSpecific(e, templates.AllGeographyTemplates(randomVariables))
}

//ExpandSpecific expands project by one entity
func (t Task) ExpandSpecific(e tools.Employee, tpls []templates.Template) error {
	for _, tpl := range tpls {
		e.RegisterJob(tpl)
	}

	return e.DoAll()
}
