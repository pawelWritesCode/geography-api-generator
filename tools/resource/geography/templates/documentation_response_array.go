package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
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
func NewDocumentationResponseArray(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("documentation_response_array").Parse(DocumentationResponseArrayTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.DocumentationDir+"response/", string(variables.Entity)+"_array.json"),
		rawTemplate, variables)
}
