//Package decay implement methods for shrinking project.
//
//To shrink project, instantiate new worker using
//	NewWorkerDecay()
//to shrink by one random entity use method
//	ShrinkRandom(picker picker.RandomEntityPicker)
package worker

import (
	"context"
	"fmt"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	templateUtils2 "generator/backend-go/tools/resource/geography/templates/templateUtils"
	picker2 "generator/backend-go/tools/resource/geography/templates/templateUtils/picker"
)

//WorkerDecay is responsible for shrinking project
type WorkerDecay struct{}

//NewWorkerDecay returns new worker struct
func NewWorkerDecay() WorkerDecay {
	return WorkerDecay{}
}

//ShrinkRandom remove one random available entity and related to it files from project
func (w WorkerDecay) ShrinkRandom(picker picker2.RandomEntityPicker) error {
	e, err := picker.RandomEntity()

	if err != nil {
		return err
	}

	return w.ShrinkSpecific(e)
}

//ShrinkSpecific entity and related to it files from project
func (w WorkerDecay) ShrinkSpecific(e templateUtils2.Entity) error {
	entityRes := resource.New(geography.EntityDir, e.EntityFU()+".php")

	if !entityRes.Exist() {
		return fmt.Errorf("entity %s does not exists", e)
	}

	resources := []resource.Resource{
		entityRes,
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

	var err error
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
