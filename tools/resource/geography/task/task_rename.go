package task

import (
	"generator/backend-go/tools"
	tUtils "generator/backend-go/tools/resource/geography/templates/templateUtils"
	gen "generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//RenameRandom rename one random available entity and related to it files
func (t Task) RenameRandom(e tools.Employee, entityGenerator gen.RandomEntity, randomPicker picker.RandomEntityAndPropertyPicker) error {
	rnd, err := randomPicker.RandomEntityAndProperty()

	if err != nil {
		return err
	}

	return t.RenameSpecificToRandom(e, rnd, entityGenerator)
}

//RenameSpecificToRandom renames specific entity for random one
func (t Task) RenameSpecificToRandom(e tools.Employee, templateVariables tUtils.TemplateVariables, entityGenerator gen.RandomEntity) error {
	return t.RenameSpecificToSpecific(e, templateVariables, entityGenerator.Random())
}

//RenameSpecificToSpecific renames specific entity for specific one
func (t Task) RenameSpecificToSpecific(employee tools.Employee, templateVariables tUtils.TemplateVariables, e tUtils.Entity) error {
	err := t.ShrinkSpecific(employee, templateVariables.Entity)

	if err != nil {
		return err
	}

	err = t.ExpandSpecific(employee, tUtils.NewTemplateVariables(e, templateVariables.Property))

	if err != nil {
		return err
	}

	return nil
}
