package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const ResourceTemplate = `AppBundle\Entity\{{.EntityFU}}:
    type: entity
    table: null
    repositoryClass: AppBundle\Repository\{{.EntityFU}}Repository
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

func NewResource(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("resource").Parse(ResourceTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: ResourcesDirectory,
		Variables: variables,
		FileName:  variables.EntityFU() + ".orm.yml",
	}
}
