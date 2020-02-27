package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const RestApiGetListTemplate = `<?php

namespace AppBundle\RestApi\{{.Entity.EntityFU}};

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
            ->entityManager->getRepository('AppBundle:{{.Entity.EntityFU}}')
            ->findBy([], ['createdAt' => 'DESC']);
    }
}
`

//NewRestApiGetList returns new template for rest api get list service
func NewRestApiGetList(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("restApiGetList").Parse(RestApiGetListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.Entity.EntityFU()+"/", "GetList.php"),
		rawTemplate, variables)
}
