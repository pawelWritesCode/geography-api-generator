package templates

import (
	"generator/backend-go/tools/generator"
	"generator/backend-go/tools/geography"
	"generator/backend-go/tools/resource"
	"log"
	"text/template"
)

const BehatDeleteTemplate = `Feature: Delete {{.Entity}} record
  As an application user
  I want to delete {{.Entity}} record

  Scenario: Delete {{.Entity}} record

    #--------------------------------------------------------------------------------
    # CREATE
    Given I generate a random string "name"
    When I send a modified "POST" request to "/api/{{.Entity}}" with data:
    """
    {
        "name": "{{"{{"}}name{{"}}"}}",
        "{{.Property}}": 3333
    }
    """
    Then the response status code should be 201
    And the response should be in JSON
    And the JSON should be valid according to schema "response/{{.Entity}}"
    And I save from the last response JSON node "id" as "{{.Entity}}Id"


    #--------------------------------------------------------------------------------
    # DELETE
    When I send a modified "DELETE" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}"
    Then the response status code should be 204


    #--------------------------------------------------------------------------------
    # DELETE
    When I send a modified "DELETE" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}"
    Then the response status code should be 404

    #--------------------------------------------------------------------------------
    # GET ONE RECORD
    When I send a modified "GET" request to "/api/{{.Entity}}/{{"{{"}}{{.Entity}}Id{{"}}"}}"
    Then the response status code should be 404`

//NewResource returns new template for resource
func NewBehatDelete(variables generator.RandomVariables) Template {
	rawTemplate, err := template.New("behat_delete").Parse(BehatDeleteTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.BehatDir+string(variables.Entity)+"/crud/", "delete.feature"),
		rawTemplate, variables)
}
