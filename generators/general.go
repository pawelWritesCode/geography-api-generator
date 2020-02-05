package generators

import (
	"errors"
	"os"
)

//ErrExpand occurs when there is no more not used entities
var ErrExpand = errors.New("every entity has been used already")

//RandomTemplateVariables returns RandomVariables type with randomly picked fields.
func RandomTemplateVariables(retries int) (RandomVariables, error)  {
	if retries == 0 {
		return RandomVariables{}, ErrExpand
	}

	entity := RandomEntity()

	//Checking if entity already exists
	_ , err := os.Stat("./backend-php/src/AppBundle/Entity/" + string(entity) + ".php")
	if !os.IsNotExist(err) {
		return RandomTemplateVariables(retries - 1)
	}

	property := RandomProperty()

	return RandomVariables{
		Entity:    entity,
		Property: property,
	}, nil
}
