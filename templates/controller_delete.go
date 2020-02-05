package templates

import (
	"generator/backend-go/generators"
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

//NewEntity returns new Template type with fulfilled fields for entity creation
func NewControllerDelete(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("controllerDelete").Parse(ControllerDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: ControllerDirectory + variables.EntityFU() + "/",
		Variables: variables,
		FileName:  "Delete.php",
	}
}
