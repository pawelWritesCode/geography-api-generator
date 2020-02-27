package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const RepositoryTemplate = `<?php

namespace AppBundle\Repository;

/**
 * {{.Entity.EntityFU}}Repository.
 *
 * This class was generated by the Doctrine ORM. Add your own custom
 * repository methods below.
 */
class {{.Entity.EntityFU}}Repository extends \Doctrine\ORM\EntityRepository
{
}
`

//NewRepository returns new template for repository
func NewRepository(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("repository").Parse(RepositoryTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RepositoryDir, variables.Entity.EntityFU()+"Repository.php"), rawTemplate, variables)
}