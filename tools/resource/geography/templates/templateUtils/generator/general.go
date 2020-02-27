//Package generator implements generating methods.
//
//To create new TemplateVariables struct required for rendering template use function:
//
//	generator.RandomTemplateVariables
//
//To generate random entity
package generator

import (
	"errors"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
)

//ErrExpand occurs when there is no more entities left to use for expanding project
var ErrExpand = errors.New("every entity has been used already")

//RandomTemplateVariables returns TemplateVariables type with randomly picked fields.
//error: ErrExpand
func RandomTemplateVariables(eGen RandomEntity, pGen RandomProperty, retries int) (templateUtils.TemplateVariables, error) {
	if retries == 0 {
		return templateUtils.TemplateVariables{}, ErrExpand
	}

	entity := eGen.Random()

	res := resource.New(geography.EntityDir, string(entity)+".php")

	if res.Exist() {
		return RandomTemplateVariables(eGen, pGen, retries-1)
	}

	property := pGen.Random()

	return templateUtils.NewTemplateVariables(entity, property), nil
}
