package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const BehatGetIdTemplate = `Feature: Get {{.Entity}} record
  As an application user
  I want to get {{.Entity}} record

  Scenario: Get {{.Entity}} record

    #--------------------------------------------------------------------------------
    # CREATE
    Given I generate a random string "name"
    When I send a modified "POST" request to "/api/{{.Entity}}" with data:
    """
    {
        "name": "{{"{{"}}name{{"}}"}}",
        "{{.Property}}": 4444
    }
    """
    Then the response status code should be 201
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And the JSON node "root.name" should be equal to templated value "{{"{{"}}name{{"}}"}}"
    And the JSON node "root.{{.Property}}" should be equal to "4444"
    And I save from the last response JSON node "id" as "{{.Entity}}Id"


    #--------------------------------------------------------------------------------
    # GET ONE RECORD
    When I send a modified "GET" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}"
    Then the response status code should be 200
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And the JSON node "root.name" should be equal to templated value "{{"{{"}}name{{"}}"}}"
    And the JSON node "root.{{.Property}}" should be equal to "4444"`

//NewResource returns new template for test get id
func NewBehatGetId(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("behat_get_id").Parse(BehatGetIdTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.BehatDir+string(variables.Entity)+"/crud/", "get_id.feature"),
		rawTemplate, variables)
}
