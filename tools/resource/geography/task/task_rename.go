package task

import (
	"generator/backend-go/tools"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	gen "generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//RenameRandomToRandom rename one random available entity and related to it files
func (t Task) RenameRandomToRandom(e tools.Employee, eGen gen.RandomEntity, rndPicker picker.RandomEntityAndPropertyPicker) error {
	randomVariables, err := rndPicker.RandomEntityAndProperty()

	if err != nil {
		return err
	}

	return t.RenameSpecificToRandom(e, randomVariables, eGen)
}

//RenameSpecificToRandom renames specific entity for random one
func (t Task) RenameSpecificToRandom(e tools.Employee, rndVariables templateUtils.TemplateVariables, entityGenerator gen.RandomEntity) error {
	renamedEntity := templateUtils.NewTemplateVariables(entityGenerator.Random(), rndVariables.Property)

	return t.RenameSpecificToSpecific(e, geography.AllGeographyResources(rndVariables.Entity), templates.AllGeographyTemplates(renamedEntity))
}

//RenameSpecificToSpecific renames specific entity for specific one
func (t Task) RenameSpecificToSpecific(employee tools.Employee, resources []resource.Resource, tpls []templates.Template) error {
	err := t.ShrinkSpecific(employee, resources)

	if err != nil {
		return err
	}

	err = t.ExpandSpecific(employee, tpls)

	if err != nil {
		return err
	}

	return nil
}
