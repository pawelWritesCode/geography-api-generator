package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const ControllerDeleteTemplate = `<?php

namespace AppBundle\Controller\{{.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\Entity\{{.EntityFU}};
use AppBundle\RestApi\{{.EntityFU}}\Delete as RestApiDeleteService;
use Symfony\Component\HttpFoundation\Response;

class Delete extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}/{{"{"}}{{.Entity}}{{"}"}}")
     * @Method("DELETE")
     */
    public function deleteAction({{.EntityFU}} ${{.Entity}}, RestApiDeleteService $service)
    {
        $service->delete(${{.Entity}});

        return $this->getJsonResponse(
            [],
            Response::HTTP_NO_CONTENT
        );
    }
}
`

//NewControllerDelete returns new template for delete controller
func NewControllerDelete(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("controllerDelete").Parse(ControllerDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ControllerDir+variables.EntityFU()+"/", "Delete.php"),
		rawTemplate, variables)
}
