package task

import (
	"fmt"
	"generator/backend-go/tools"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//ShrinkRandom remove one random available entity and related to it files from project
func (t Task) ShrinkRandom(employee tools.Employee, picker picker.RandomEntityPicker) error {
	e, err := picker.RandomEntity()

	if err != nil {
		return err
	}

	return t.ShrinkSpecific(employee, geography.AllGeographyResources(e))
}

//ShrinkSpecific entity and related to it files from project
func (t Task) ShrinkSpecific(employee tools.Employee, resources []resource.Resource) error {
	for _, res := range resources {
		if !res.Exist() {
			return fmt.Errorf("resource %s does not exists", res)
		}
	}

	for _, res := range resources {
		employee.RegisterJob(res)
	}

	return employee.DoAll()
}
