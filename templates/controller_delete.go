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

//NewControllerDelete returns new template for delete controller
func NewControllerDelete(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("controllerDelete").Parse(ControllerDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: ControllerDirectory + variables.EntityFU() + "/",
		FileName:  "Delete.php",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
