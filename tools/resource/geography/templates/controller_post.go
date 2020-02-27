package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const ControllerPostTemplate = `<?php

namespace AppBundle\Controller\{{.Entity.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\RestApi\{{.Entity.EntityFU}}\Post as RestApiPostService;
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
func NewControllerPost(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("controllerPost").Parse(ControllerPostTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ControllerDir+variables.Entity.EntityFU()+"/", "Post.php"),
		rawTemplate, variables)
}
