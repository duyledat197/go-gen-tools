{{define "delete"}}
package features

import "github.com/cucumber/godog"

func {{.CamelCase}}HaveBeenDeletedCorrectly() error {
       return godog.ErrPending
}

func userDelete{{.PascalCase}}() error {
       return godog.ErrPending
}

func userDelete{{.PascalCase}}Again() error {
       return godog.ErrPending
}
{{end}}