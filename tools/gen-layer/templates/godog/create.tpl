{{define "create"}}
package features

import "github.com/cucumber/godog"

func {{.CamelCase}}MustBeCreated() error {
       return godog.ErrPending
}

func userCreate{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}