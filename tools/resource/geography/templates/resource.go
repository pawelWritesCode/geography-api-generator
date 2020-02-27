package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const ResourceTemplate = `AppBundle\Entity\{{.Entity.EntityFU}}:
    type: entity
    table: null
    repositoryClass: AppBundle\Repository\{{.Entity.EntityFU}}Repository
    id:
        id:
            type: integer
            id: true
            generator:
                strategy: AUTO
    fields:
        name:
            type: string
            length: 255
        {{.Property}}:
            type: integer
        createdAt:
            type: datetime
            nullable: true
    lifecycleCallbacks:
        prePersist:
            - updateTimestamps
`

//NewResource returns new template for resource
func NewResource(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("resource").Parse(ResourceTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ResourcesDir, variables.Entity.EntityFU()+".orm.yml"), rawTemplate, variables)
}
