package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const RestApiGetListTemplate = `<?php

namespace AppBundle\RestApi\{{.EntityFU}};

use Doctrine\ORM\EntityManagerInterface;

class GetList
{
    private $entityManager;

    public function __construct(EntityManagerInterface $entityManager)
    {
        $this->entityManager = $entityManager;
    }

    public function get()
    {
        return $this
            ->entityManager->getRepository('AppBundle:{{.EntityFU}}')
            ->findBy([], ['createdAt' => 'DESC']);
    }
}
`

//NewRestApiGetList returns new template for rest api get list service
func NewRestApiGetList(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("restApiGetList").Parse(RestApiGetListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: RestApiDirectory + variables.EntityFU() + "/",
		FileName:  "GetList.php",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
