package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const RestApiPostTemplate = `<?php

namespace AppBundle\RestApi\{{.Entity.EntityFU}};

use AppBundle\Entity\{{.Entity.EntityFU}};
use Doctrine\ORM\EntityManagerInterface;

class Post
{
    private $entityManager;

    public function __construct(
        EntityManagerInterface $entityManager
    ) {
        $this->entityManager = $entityManager;
    }

    public function post($requestData)
    {
        ${{.Entity}} = $this->entityManager
            ->getRepository('AppBundle:{{.Entity.EntityFU}}')
            ->findOneBy(['name' => $requestData['name']]);

        if (!${{.Entity}}) {
            ${{.Entity}} = new {{.Entity.EntityFU}}();
            ${{.Entity}}->setFromArray($requestData);
            $this->entityManager->persist(${{.Entity}});
            $this->entityManager->flush();
        }

        return ${{.Entity}};
    }
}
`

//NewRestApiPost returns new template for rest api post service
func NewRestApiPost(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("restApiPost").Parse(RestApiPostTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.RestApiDir+variables.Entity.EntityFU()+"/",
		"Post.php"), rawTemplate, variables)
}
