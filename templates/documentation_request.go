package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const DocumentationRequestTemplate = `{
    "title": "Request: {{.Entity}}",
    "type": "object",
    "properties": {
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
    "additionalProperties": false
}
`

//NewDocumentationRequest returns new template for documentation request
func NewDocumentationRequest(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("documentation_request").Parse(DocumentationRequestTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: DocumentationDirectory + "request/",
		FileName:  string(variables.Entity) + ".json",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
