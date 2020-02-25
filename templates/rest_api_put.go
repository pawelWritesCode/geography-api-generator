package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const RestApiPutTemplate = `<?php

namespace AppBundle\RestApi\{{.Entity.EntityFU}};

use AppBundle\Entity\{{.Entity.EntityFU}};
use Doctrine\ORM\EntityManagerInterface;

class Put
{
    private $entityManager;

    public function __construct(EntityManagerInterface $entityManager)
    {
        $this->entityManager = $entityManager;
    }

    /**
     * @param {{.Entity.EntityFU}} ${{.Entity}}
     * @param $requestData
     */
    public function put({{.Entity.EntityFU}} ${{.Entity}}, $requestData)
    {
        ${{.Entity}}->setFromArray($requestData);

        $this->entityManager->flush();
    }
}
`

//NewRestApiPut returns new template for rest api put service
func NewRestApiPut(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("restApiPut").Parse(RestApiPutTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.Entity.EntityFU()+"/", "Put.php"),
		rawTemplate, variables)
}
