{{define "list"}}Feature: list {{.CamelCase}}

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when list {{.CamelCase}}
        Given a signed in "<role>"
        When user list {{.CamelCase}}
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# list {{.CamelCase}}
    Scenario: list {{.CamelCase}}
        When user list {{.CamelCase}}
        Then returns "OK" status code
        And our system must return results correctly
{{end}}