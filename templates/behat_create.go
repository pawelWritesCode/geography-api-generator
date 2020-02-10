package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const BehatCreateTemplate = `Feature: Create {{.Entity}} record
  As an application user
  I want to create {{.Entity}} record

  Scenario: Create {{.Entity}} record

    #--------------------------------------------------------------------------------
    # CREATE
    Given I generate a random string "name"
    When I send a modified "POST" request to "/api/{{.Entity}}" with data:
    """
    {
        "name": "{{"{{"}}name{{"}}"}}",
        "{{.Property}}": 1111
    }
    """
    Then the response status code should be 201
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And the JSON node "root.name" should be equal to templated value "{{"{{"}}name{{"}}"}}"
    And the JSON node "root.{{.Property}}" should be equal to "1111"`

//NewResource returns new template for resource
func NewBehatCreate(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("behat_create").Parse(BehatCreateTemplate)

	if err != nil {
		log.Fatal(err)
	}

	resource := Resource{
		Directory: BehatDirectory + string(variables.Entity) + "/crud/",
		FileName:  "create.feature",
	}

	return Template{
		Payload:   rawTemplate,
		Variables: variables,
		Resource:  resource,
	}
}
