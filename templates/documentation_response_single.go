package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
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
func NewDocumentationResponseSingle(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("documentation_response_single").Parse(DocumentationResponseSingleTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.DocumentationDir+"response/", string(variables.Entity)+".json"),
		rawTemplate, variables)
}
