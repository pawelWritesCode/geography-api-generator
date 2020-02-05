package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const RestApiPutTemplate = `<?php

namespace AppBundle\RestApi\{{.EntityFU}};

use AppBundle\Entity\{{.EntityFU}};
use Doctrine\ORM\EntityManagerInterface;

class Put
{
    private $entityManager;

    public function __construct(EntityManagerInterface $entityManager)
    {
        $this->entityManager = $entityManager;
    }

    /**
     * @param {{.EntityFU}} ${{.Entity}}
     * @param $requestData
     */
    public function put({{.EntityFU}} ${{.Entity}}, $requestData)
    {
        ${{.Entity}}->setFromArray($requestData);

        $this->entityManager->flush();
    }
}
`

//NewRestApiPut returns new template for rest api put service
func NewRestApiPut(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("restApiPut").Parse(RestApiPutTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: RestApiDirectory + variables.EntityFU() + "/",
		Variables: variables,
		FileName:  "Put.php",
	}
}
