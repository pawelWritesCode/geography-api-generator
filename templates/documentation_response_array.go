package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const DocumentationResponseArrayTemplate = `{
    "title": "Response: {{.Entity}} array",
    "type": "array",
    "items": {
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
}
`

//NewDocumentationResponseArray returns new template for documentation response array
func NewDocumentationResponseArray(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("documentation_response_array").Parse(DocumentationResponseArrayTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: DocumentationDirectory + "response/",
		FileName:  string(variables.Entity) + "_array.json",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
