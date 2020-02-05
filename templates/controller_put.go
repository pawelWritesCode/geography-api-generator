package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const ControllerPutTemplate = `<?php

namespace AppBundle\Controller\{{.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\Entity\{{.EntityFU}};
use AppBundle\RestApi\{{.EntityFU}}\Put as RestApiPutService;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Serializer\SerializerInterface;

class Put extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}/{{"{"}}{{.Entity}}{{"}"}}")
     * @Method("PUT")
     */
    public function putAction(
        Request $request,
        SerializerInterface $serializer,
        {{.EntityFU}} ${{.Entity}},
        RestApiPutService $service
    ) {
        $jsonRequestData = json_decode($request->getContent());

        $this->validateDataAgainstSchema(
            '{{.Entity}}',
            $jsonRequestData,
            Response::HTTP_FORBIDDEN
        );

        $requestData = (array) $jsonRequestData;

        $service->put(${{.Entity}}, $requestData);

        $data = $serializer->serialize(
            ${{.Entity}},
            'json',
            ['groups' => ['default']]
        );

        return $this->getJsonResponse(
            $data
        );
    }
}
`

//NewControllerPut returns new template for put controller
func NewControllerPut(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("controllerPut").Parse(ControllerPutTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: ControllerDirectory + variables.EntityFU() + "/",
		Variables: variables,
		FileName:  "Put.php",
	}
}
