//Package picker implement method for picking project artifacts
package picker

import (
	"bufio"
	"errors"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//Picker is object responsible for picking entity related things
type Picker struct{}

type RandomEntityPicker interface {
	RandomEntity() (templateUtils.Entity, error)
}

type RandomEntityAndPropertyPicker interface {
	RandomEntityAndProperty() (templateUtils.TemplateVariables, error)
}

//ErrNoAvailableEntities occurs when there are no left any free entities to shrink from
var ErrNoAvailableEntities = errors.New("no available entities")

//NewTemplateVariables returns new Picker struct
func New() Picker {
	return Picker{}
}

//RandomEntity pick one of available entities
func (p Picker) RandomEntity() (templateUtils.Entity, error) {
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

//RandomEntityAndProperty pick one of available entities and returns templateVariables with entity and property
func (p Picker) RandomEntityAndProperty() (templateUtils.TemplateVariables, error) {
	entity, err := p.RandomEntity()
	randomVariable := templateUtils.NewTemplateVariables(entity, "")

	if err != nil {
		return randomVariable, err
	}

	return p.EntityAndProperty(entity)
}

//EntityAndProperty returns templateVariables with entity and its property
func (p Picker) EntityAndProperty(e templateUtils.Entity) (templateUtils.TemplateVariables, error) {
	randomVariable := templateUtils.NewTemplateVariables(e, "")
	readFile, err := os.Open(geography.EntityDir + e.EntityFU() + ".php")

	if err != nil {
		return randomVariable, err
	}

	wordsToOmit := []string{"$id;", "$name;", "$createdAt;"}
	scanner := bufio.NewScanner(readFile)

	var propertyRaw string
	shouldContinue := true
	for scanner.Scan() && shouldContinue {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		words := strings.Split(trimmedLine, " ")

		if len(words) == 2 && words[0] == "private" && !inArray(words[1], wordsToOmit) {
			shouldContinue = false
			propertyRaw = words[1]
		}
	}

	if err = scanner.Err(); err != nil {
		return randomVariable, err
	}

	randomVariable.Property = templateUtils.Property(propertyRaw[1 : len(propertyRaw)-1])

	return randomVariable, err
}

//AvailableEntities loops over entity folder and returns array of available entities
func (p Picker) AvailableEntities() ([]templateUtils.Entity, error) {
	entities := []templateUtils.Entity{}
	err := filepath.Walk(geography.EntityDir, func(path string, info os.FileInfo, err error) error {
		entityWithExtension := filepath.Base(path)
		entityName := strings.TrimSuffix(entityWithExtension, ".php")

		if entityName == "EntityExtensions" || info.IsDir() {
			return nil
		}

		entities = append(entities, templateUtils.Entity(strings.ToLower(entityName)))

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
func (p Picker) EntityExists(e templateUtils.Entity) (bool, error) {
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

//inArray checks if item is in given arr
func inArray(item string, arr []string) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}

	return false
}
