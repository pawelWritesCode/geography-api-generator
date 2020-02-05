package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const RepositoryTemplate = `<?php

namespace AppBundle\Repository;

/**
 * {{.EntityFU}}Repository.
 *
 * This class was generated by the Doctrine ORM. Add your own custom
 * repository methods below.
 */
class {{.EntityFU}}Repository extends \Doctrine\ORM\EntityRepository
{
}
`

func NewRepository(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("repository").Parse(RepositoryTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: RepositoryDirectory,
		Variables: variables,
		FileName:  variables.EntityFU() + "Repository.php",
	}
}
