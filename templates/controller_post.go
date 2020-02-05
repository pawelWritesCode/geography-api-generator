package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const ControllerPostTemplate = `<?php

namespace AppBundle\Controller\{{.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\RestApi\{{.EntityFU}}\Post as RestApiPostService;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Serializer\SerializerInterface;

class Post extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}")
     * @Method({"POST"})
     */
    public function postAction(Request $request, SerializerInterface $serializer, RestApiPostService $service)
    {
        $jsonRequestData = json_decode($request->getContent());

        $this->validateDataAgainstSchema(
            '{{.Entity}}',
            $jsonRequestData,
            Response::HTTP_FORBIDDEN
        );

        $requestData = (array) $jsonRequestData;

        ${{.Entity}} = $service->post($requestData);

        $data = $serializer->serialize(
            ${{.Entity}},
            'json',
            ['groups' => ['default']]
        );

        return $this->getJsonResponse(
            $data,
            Response::HTTP_CREATED
        );
    }
}
`

//NewControllerPost returns new template for post controller
func NewControllerPost(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("controllerPost").Parse(ControllerPostTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: ControllerDirectory + variables.EntityFU() + "/",
		Variables: variables,
		FileName:  "Post.php",
	}
}
