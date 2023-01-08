{{define "retrieve"}}Feature: retrieve {{.CamelCase}}

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when retrieve {{.CamelCase}}
        Given a signed in "<role>"
        When user retrieve {{.CamelCase}}
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# retrieve {{.CamelCase}}
    Scenario: retrieve {{.CamelCase}}
        When user retrieve {{.CamelCase}}
        Then returns "OK" status code
        And our system must return result correctly
	
	# retrieve invalid {{.CamelCase}}
    Scenario: retrieve invalid {{.CamelCase}}
        Given {{.CamelCase}} is deleted
        When user retrieve {{.CamelCase}}
        Then returns "NotFound" status code
{{end}}