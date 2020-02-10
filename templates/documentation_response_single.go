package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const DocumentationResponseSingleTemplate = `{
    "title": "Response: {{.Entity}}",
    "type": "object",
    "properties": {
        "id": {
            "type": "integer"
        },
        "name": {
            "type": "string"
        },
        "{{.Property}}": {
            "type": "integer"
        },
        "createdAt": [
            "string",
            "null"
        ]
    },
    "required": [
        "id",
        "name",
        "{{.Property}}",
        "createdAt"
    ],
    "additionalProperties": false
}
`

//NewDocumentationResponseSingle returns new template for documentation response single
func NewDocumentationResponseSingle(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("documentation_response_single").Parse(DocumentationResponseSingleTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: DocumentationDirectory + "response/",
		FileName:  string(variables.Entity) + ".json",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
