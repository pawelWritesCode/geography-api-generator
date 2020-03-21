package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const BehatPutTemplate = `Feature: Update {{.Entity}} record
  As an application user
  I want to Update {{.Entity}} record

  Scenario: Update {{.Entity}} record

    #--------------------------------------------------------------------------------
    # CREATE
    Given I generate a random string "name"
    When I send a modified "POST" request to "/api/{{.Entity}}" with data:
    """
    {
        "name": "{{"{{"}}name{{"}}"}}",
        "{{.Property}}": 2222
    }
    """
    Then the response status code should be 201
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And I save from the last response JSON node "id" as "{{.Entity}}Id"


    #--------------------------------------------------------------------------------
    # UPDATE
    Given I generate a random string "name2"
    When I send a modified "PUT" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}" with data:
    """
    {
        "name": "{{"{{"}}name2{{"}}"}}",
        "{{.Property}}": 99999
    }
    """
    Then the response status code should be 200
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And the JSON node "root.name" should be equal to templated value "{{"{{"}}name2{{"}}"}}"
    And the JSON node "root.{{.Property}}" should be equal to "99999"


    #--------------------------------------------------------------------------------
    # GET ONE RECORD
    When I send a modified "GET" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}"
    Then the response status code should be 200
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And the JSON node "root.name" should be equal to templated value "{{"{{"}}name2{{"}}"}}"
    And the JSON node "root.{{.Property}}" should be equal to "99999"`

//NewResource returns new template for test put
func NewBehatPut(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("behat_put").Parse(BehatPutTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.BehatDir+string(variables.Entity)+"/crud/", "put.feature"), rawTemplate, variables)
}
