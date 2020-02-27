package worker

import (
	tUtils "generator/backend-go/tools/resource/geography/templates/templateUtils"
	gen "generator/backend-go/tools/resource/geography/templates/templateUtils/generator"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//WorkerRename is responsible for shrinking project
type WorkerRename struct{}

//NewWorkerRename returns new worker struct
func NewWorkerRename() WorkerRename {
	return WorkerRename{}
}

//RenameRandom rename one random available entity and related to it files
func (w WorkerRename) RenameRandom(entityGenerator gen.RandomEntity, randomPicker picker.RandomEntityAndPropertyPicker) error {
	rnd, err := randomPicker.RandomEntityAndProperty()

	if err != nil {
		return err
	}

	return w.RenameSpecificToRandom(rnd, entityGenerator)
}

//RenameSpecificToRandom renames specific entity for random one
func (w WorkerRename) RenameSpecificToRandom(templateVariables tUtils.TemplateVariables, entityGenerator gen.RandomEntity) error {
	return w.RenameSpecificToSpecific(templateVariables, entityGenerator.Random())
}

//RenameSpecificToSpecific renames specific entity for specific one
func (w WorkerRename) RenameSpecificToSpecific(templateVariables tUtils.TemplateVariables, e tUtils.Entity) error {
	decayWorker := NewWorkerDecay()
	err := decayWorker.ShrinkSpecific(templateVariables.Entity)

	if err != nil {
		return err
	}

	specificVariables := tUtils.NewTemplateVariables(e, templateVariables.Property)

	expandWorker := NewWorkerExpand()
	err = expandWorker.ExpandSpecific(specificVariables)

	if err != nil {
		return err
	}

	return nil
}
