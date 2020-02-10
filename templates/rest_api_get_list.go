package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
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
func NewRestApiGetList(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("restApiGetList").Parse(RestApiGetListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.EntityFU()+"/", "GetList.php"),
		rawTemplate, variables)
}
