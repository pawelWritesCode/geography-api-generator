package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const ControllerGetTemplate = `<?php

namespace AppBundle\Controller\{{.Entity.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\Entity\{{.Entity.EntityFU}};
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Serializer\SerializerInterface;

class Get extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}/{{"{"}}{{.Entity}}{{"}"}}")
     * @Method("GET")
     */
    public function getAction(
        Request $request,
        SerializerInterface $serializer,
        {{.Entity.EntityFU}} ${{.Entity}}
    ) {
        $data = $serializer->serialize(
            ${{.Entity}},
            'json',
            ['groups' => ['default']]
        );

        return $this->getDebugOrJsonResponse(
            $request,
            $data
        );
    }
}
`

//NewControllerGet returns new template for get controller
func NewControllerGet(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("controllerGet").Parse(ControllerGetTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ControllerDir+variables.Entity.EntityFU()+"/", "Get.php"),
		rawTemplate, variables)
}
