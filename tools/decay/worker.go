//Package decay implement method for shrinking project
package decay

import (
	"fmt"
	"generator/backend-go/tools/decay/picker"
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
)

//Worker is responsible for shrinking project
type Worker struct{}

//Shrink remove entity related files from project
func (w Worker) Shrink(e generator.Entity) error {
	entityPicker := picker.New()
	ok, err := entityPicker.EntityExists(e)

	if err != nil {
		return err
	}

	if ok == false {
		return fmt.Errorf("entity %s does not exists", e)
	}

	resources := []resource.Resource{
		resource.New(geography.EntityDir, e.EntityFU()+".php"),
		resource.New(geography.RepositoryDir, e.EntityFU()+"Repository.php"),
		resource.New(geography.ResourcesDir, e.EntityFU()+".orm.yml"),
		resource.New(geography.ControllerDir+e.EntityFU()+"/", ""),
		resource.New(geography.RestApiDir+e.EntityFU()+"/", ""),
		resource.New(geography.BehatDir+string(e)+"/", ""),
		resource.New(geography.DocumentationDir+"request/", string(e)+".json"),
		resource.New(geography.DocumentationDir+"response/", string(e)+".json"),
		resource.New(geography.DocumentationDir+"response/", string(e)+"_array.json"),
	}

	for _, res := range resources {
		err := res.Unlink()

		if err != nil {
			return err
		}
	}

	return nil
}
