package picker

import (
	"errors"
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//Picker is object responsible for picking entity related things
type Picker struct{}

//ErrNoAvailableEntities occurs when there are no left any free entities to shrink from
var ErrNoAvailableEntities = errors.New("no available entities")

//New returns new Picker struct
func New() Picker {
	return Picker{}
}

//RandomEntity pick one of available entities
func (p Picker) RandomEntity() (generator.Entity, error) {
	entities, err := p.AvailableEntities()
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(entities) - 1
	randIndex := rand.Intn(max-min+1) + min

	return entities[randIndex], nil
}

//AvailableEntities loops over entity folder and returns array of available entities
func (p Picker) AvailableEntities() ([]generator.Entity, error) {
	entities := []generator.Entity{}
	err := filepath.Walk(geography.EntityDir, func(path string, info os.FileInfo, err error) error {
		entityWithExtension := filepath.Base(path)
		entityName := strings.TrimSuffix(entityWithExtension, ".php")

		if entityName == "EntityExtensions" || info.IsDir() {
			return nil
		}
		entities = append(entities, generator.Entity(entityName))

		return nil
	})

	if err != nil {
		return entities, err
	}

	if len(entities) == 0 {
		return entities, ErrNoAvailableEntities
	}

	return entities, nil
}

//EntityExists checks if given entity is available
func (p Picker) EntityExists(e generator.Entity) (bool, error) {
	entities, err := p.AvailableEntities()

	if err != nil {
		return false, err
	}

	for _, entity := range entities {
		if entity == e {
			return true, nil
		}
	}

	return false, nil
}
