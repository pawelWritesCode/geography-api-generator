package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const RestApiDeleteTemplate = `<?php

namespace AppBundle\RestApi\{{.EntityFU}};

use AppBundle\Entity\{{.EntityFU}};
use Doctrine\ORM\EntityManagerInterface;

class Delete
{
    private $entityManager;

    public function __construct(
        EntityManagerInterface $entityManager
    ) {
        $this->entityManager = $entityManager;
    }

    public function delete({{.EntityFU}} ${{.Entity}})
    {
        $this->entityManager->remove(${{.Entity}});
        $this->entityManager->flush();
    }
}
`

//NewRestApiDelete returns new template for rest api delete service
func NewRestApiDelete(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("restApiDelete").Parse(RestApiDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.EntityFU()+"/", "Delete.php"),
		rawTemplate, variables)
}
