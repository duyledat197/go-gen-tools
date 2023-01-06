{{define "update"}}Feature: update {{.CamelCase}}

    Background: basic background
        Given a signed in "admin"
		And a background

	# authenticate
    Scenario Outline: authenticate when update {{.CamelCase}}
        Given a signed in "<role>"
        When user update {{.CamelCase}}
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# update {{.CamelCase}}
    Scenario: update {{.CamelCase}}
        When user update {{.CamelCase}}
        Then returns "OK" status code
        And updated {{.CamelCase}} set as expected

	# update invalid {{.CamelCase}}
    Scenario: update invalid {{.CamelCase}}
        Given {{.CamelCase}} is deleted
        When user update {{.CamelCase}}
        Then returns "NotFound" status code
{{end}}