{{define "create"}}Feature: Create {{.CamelCase}}

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when create {{.CamelCase}}
        Given a signed in "<role>"
        When user create {{.CamelCase}}
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# create {{.CamelCase}}
    Scenario: create {{.CamelCase}}
        When user create {{.CamelCase}}
        Then returns "OK" status code
        And {{.CamelCase}} must be created
{{end}}