package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const RestApiDeleteTemplate = `<?php

namespace AppBundle\RestApi\{{.Entity.EntityFU}};

use AppBundle\Entity\{{.Entity.EntityFU}};
use Doctrine\ORM\EntityManagerInterface;

class Delete
{
    private $entityManager;

    public function __construct(
        EntityManagerInterface $entityManager
    ) {
        $this->entityManager = $entityManager;
    }

    public function delete({{.Entity.EntityFU}} ${{.Entity}})
    {
        $this->entityManager->remove(${{.Entity}});
        $this->entityManager->flush();
    }
}
`

//NewRestApiDelete returns new template for rest api delete service
func NewRestApiDelete(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("restApiDelete").Parse(RestApiDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.Entity.EntityFU()+"/", "Delete.php"),
		rawTemplate, variables)
}
