package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const ControllerGetListTemplate = `<?php

namespace AppBundle\Controller\{{.Entity.EntityFU}};

use Sensio\Bundle\FrameworkExtraBundle\Configuration\Method;
use Sensio\Bundle\FrameworkExtraBundle\Configuration\Route;
use AppBundle\Controller\GenericController;
use AppBundle\RestApi\{{.Entity.EntityFU}}\GetList as RestApiGetListService;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Serializer\SerializerInterface;

class GetList extends GenericController
{
    /**
     * @Route("/api/{{.Entity}}")
     * @Method({"GET"})
     */
    public function getList(
        Request $request,
        SerializerInterface $serializer,
        RestApiGetListService $service
    ) {
        $items = $service->get();

        $data = $serializer->serialize(
            $items,
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

//NewControllerGetList returns new template for get list controller
func NewControllerGetList(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("controllerGetList").Parse(ControllerGetListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ControllerDir+variables.Entity.EntityFU()+"/", "GetList.php"),
		rawTemplate, variables)
}
