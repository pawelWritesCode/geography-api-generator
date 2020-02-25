package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
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
func NewResource(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("resource").Parse(ResourceTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.ResourcesDir, variables.Entity.EntityFU()+".orm.yml"), rawTemplate, variables)
}
