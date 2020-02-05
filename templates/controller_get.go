package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const ControllerGetTemplate = `<?php

namespace AppBundle\Controller\{{.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\Entity\{{.EntityFU}};
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
        {{.EntityFU}} ${{.Entity}}
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
func NewControllerGet(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("controllerGet").Parse(ControllerGetTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: ControllerDirectory + variables.EntityFU() + "/",
		Variables: variables,
		FileName:  "Get.php",
	}
}
