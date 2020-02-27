package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const ControllerDeleteTemplate = `<?php

namespace AppBundle\Controller\{{.Entity.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\Entity\{{.Entity.EntityFU}};
use AppBundle\RestApi\{{.Entity.EntityFU}}\Delete as RestApiDeleteService;
use Symfony\Component\HttpFoundation\Response;

class Delete extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}/{{"{"}}{{.Entity}}{{"}"}}")
     * @Method("DELETE")
     */
    public function deleteAction({{.Entity.EntityFU}} ${{.Entity}}, RestApiDeleteService $service)
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
func NewControllerDelete(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("controllerDelete").Parse(ControllerDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ControllerDir+variables.Entity.EntityFU()+"/", "Delete.php"),
		rawTemplate, variables)
}
