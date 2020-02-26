//Package decay implement methods for shrinking project.
//
//To shrink project, instantiate new worker using
//	New()
//to shrink by one random entity use method
//	ShrinkRandom(picker picker.RandomEntityPicker)
package decay

import (
	"context"
	"generator/backend-go/tools/decay/picker"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
)

//Worker is responsible for shrinking project
type Worker struct{}

//New returns new worker struct
func New() Worker {
	return Worker{}
}

//ShrinkRandom remove random available entity and related to it files from project
func (w Worker) ShrinkRandom(picker picker.RandomEntityPicker) error {
	e, err := picker.RandomEntity()

	if err != nil {
		return err
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

	ch1 := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, res := range resources {
		go unlinkResource(ctx, res, ch1)
	}

	for i := 0; i < len(resources); i++ {
		err = <-ch1

		if err != nil {
			cancel()
			return err
		}
	}

	return nil
}

//unlinkResource unlinks resource
func unlinkResource(ctx context.Context, res resource.Resource, ch1 chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		ch1 <- res.Unlink()
	}
}
