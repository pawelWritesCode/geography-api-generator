package task

import (
	"fmt"
	"generator/backend-go/tools"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//ShrinkRandom remove one random available entity and related to it files from project
func (t Task) ShrinkRandom(employee tools.Employee, picker picker.RandomEntityPicker) error {
	e, err := picker.RandomEntity()

	if err != nil {
		return err
	}

	return t.ShrinkSpecific(employee, e)
}

//ShrinkSpecific entity and related to it files from project
func (t Task) ShrinkSpecific(employee tools.Employee, e templateUtils.Entity) error {
	entityRes := resource.New(geography.EntityDir, e.EntityFU()+".php")

	if !entityRes.Exist() {
		return fmt.Errorf("entity %s does not exists", e)
	}

	employee.RegisterJob(entityRes)
	employee.RegisterJob(resource.New(geography.RepositoryDir, e.EntityFU()+"Repository.php"))
	employee.RegisterJob(resource.New(geography.ResourcesDir, e.EntityFU()+".orm.yml"))
	employee.RegisterJob(resource.New(geography.ControllerDir+e.EntityFU()+"/", ""))
	employee.RegisterJob(resource.New(geography.RestApiDir+e.EntityFU()+"/", ""))
	employee.RegisterJob(resource.New(geography.BehatDir+string(e)+"/", ""))
	employee.RegisterJob(resource.New(geography.DocumentationDir+"request/", string(e)+".json"))
	employee.RegisterJob(resource.New(geography.DocumentationDir+"response/", string(e)+".json"))
	employee.RegisterJob(resource.New(geography.DocumentationDir+"response/", string(e)+"_array.json"))

	return employee.DoAll()
}
