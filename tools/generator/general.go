package generator

import (
	"errors"
	"generator/backend-go/tools/geography"
	resource3 "generator/backend-go/tools/resource"
)

//ErrExpand occurs when there is no more entities left to use for expanding project
var ErrExpand = errors.New("every entity has been used already")

//RandomTemplateVariables returns RandomVariables type with randomly picked fields.
//error: ErrExpand
func RandomTemplateVariables(eGen RandomEntity, pGen RandomProperty, retries int) (RandomVariables, error) {
	if retries == 0 {
		return RandomVariables{}, ErrExpand
	}

	entity := eGen.Random()

	resource := resource3.Resource{
		Directory: geography.EntityDir,
		FileName:  string(entity) + ".php",
	}

	if resource.Exist() {
		return RandomTemplateVariables(eGen, pGen, retries-1)
	}

	property := pGen.Random()

	return RandomVariables{
		Entity:   entity,
		Property: property,
	}, nil
}
