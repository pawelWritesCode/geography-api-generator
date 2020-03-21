package task

import (
	tUtils "generator/backend-go/tools/resource/geography/templates/templateUtils"
	gen "generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//RenameRandom rename one random available entity and related to it files
func (t Task) RenameRandom(entityGenerator gen.RandomEntity, randomPicker picker.RandomEntityAndPropertyPicker) error {
	rnd, err := randomPicker.RandomEntityAndProperty()

	if err != nil {
		return err
	}

	return t.RenameSpecificToRandom(rnd, entityGenerator)
}

//RenameSpecificToRandom renames specific entity for random one
func (t Task) RenameSpecificToRandom(templateVariables tUtils.TemplateVariables, entityGenerator gen.RandomEntity) error {
	return t.RenameSpecificToSpecific(templateVariables, entityGenerator.Random())
}

//RenameSpecificToSpecific renames specific entity for specific one
func (t Task) RenameSpecificToSpecific(templateVariables tUtils.TemplateVariables, e tUtils.Entity) error {
	err := t.ShrinkSpecific(templateVariables.Entity)

	if err != nil {
		return err
	}

	err = t.ExpandSpecific(tUtils.NewTemplateVariables(e, templateVariables.Property))

	if err != nil {
		return err
	}

	return nil
}
