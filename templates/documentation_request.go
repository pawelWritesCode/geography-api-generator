package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
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
func NewDocumentationRequest(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("documentation_request").Parse(DocumentationRequestTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.DocumentationDir+"request/", string(variables.Entity)+".json"),
		rawTemplate, variables)
}
