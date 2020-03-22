//Package task implements tasks responsible for expanding, shrinking and renaming geography.
//
// Instantiating
//	func New()
//
// Expanding
//
//To expand project use methods starting with Expand
//	func (t Task) ExpandRandom(eGen generator.RandomEntity, pGen generator.RandomProperty) error
//	func (t Task) ExpandSpecific(randomVariables templateUtils.TemplateVariables) error
//
// Shrinking
//
//To shrink project use methods starting with Shrink.
//	func (t Task) ShrinkRandom(picker picker.RandomEntityPicker) error
//	func (t Task) ShrinkSpecific(e templateUtils.Entity) error
//
// Renaming
//
//To rename project entity use methods starting with Rename
//
//	func (t Task) RenameRandomToRandom(entityGenerator gen.RandomEntity, randomPicker picker.RandomEntityAndPropertyPicker) error
//	func (t Task) RenameSpecificToRandom(templateVariables tUtils.TemplateVariables, entityGenerator gen.RandomEntity) error
//	func (t Task) RenameSpecificToSpecific(templateVariables tUtils.TemplateVariables, e tUtils.Entity) error
package task

//Task represents some kind of task to do
type Task struct{}

//New returns new Task instance
func New() Task {
	return Task{}
}
