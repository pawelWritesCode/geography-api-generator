package task

import (
	"fmt"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
	"generator/backend-go/tools/resource/geography/worker"
)

//ShrinkRandom remove one random available entity and related to it files from project
func (t Task) ShrinkRandom(picker picker.RandomEntityPicker) error {
	e, err := picker.RandomEntity()

	if err != nil {
		return err
	}

	return t.ShrinkSpecific(e)
}

//ShrinkSpecific entity and related to it files from project
func (t Task) ShrinkSpecific(e templateUtils.Entity) error {
	entityRes := resource.New(geography.EntityDir, e.EntityFU()+".php")

	if !entityRes.Exist() {
		return fmt.Errorf("entity %s does not exists", e)
	}

	w := worker.NewWorker()

	w.RegisterJob(entityRes)
	w.RegisterJob(resource.New(geography.RepositoryDir, e.EntityFU()+"Repository.php"))
	w.RegisterJob(resource.New(geography.ResourcesDir, e.EntityFU()+".orm.yml"))
	w.RegisterJob(resource.New(geography.ControllerDir+e.EntityFU()+"/", ""))
	w.RegisterJob(resource.New(geography.RestApiDir+e.EntityFU()+"/", ""))
	w.RegisterJob(resource.New(geography.BehatDir+string(e)+"/", ""))
	w.RegisterJob(resource.New(geography.DocumentationDir+"request/", string(e)+".json"))
	w.RegisterJob(resource.New(geography.DocumentationDir+"response/", string(e)+".json"))
	w.RegisterJob(resource.New(geography.DocumentationDir+"response/", string(e)+"_array.json"))

	return w.DoAll()
}
