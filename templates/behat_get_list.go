package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const BehatGetListTemplate = `Feature: Get list {{.Entity}} record
  As an application user
  I want to get list {{.Entity}} record

  Scenario: Get list {{.Entity}} record

    #--------------------------------------------------------------------------------
    # CREATE
    Given I generate a random string "name"
    When I send a modified "POST" request to "/api/{{.Entity}}" with data:
    """
    {
        "name": "{{"{{"}}name{{"}}"}}",
        "{{.Property}}": 5555
    }
    """
    Then the response status code should be 201
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And I save from the last response JSON node "id" as "{{.Entity}}Id"

    #--------------------------------------------------------------------------------
    # GET LIST
    When I send a modified "GET" request to "/api/{{.Entity}}"
    Then the response status code should be 200
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}_array"
    And the JSON should be valid according to this schema:
    """
      {
          "type": "array",
          "minItems": 1
      }
      """
    And list element with the id "{{"{{"}}{{.Entity}}Id{{"}}"}}" has field "name" with value "{{"{{"}}name{{"}}"}}"`

//NewResource returns new template for resource
func NewBehatGetList(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("behat_get_list").Parse(BehatGetListTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: BehatDirectory + string(variables.Entity) + "/crud/",
		Variables: variables,
		FileName:  "get_list.feature",
	}
}
