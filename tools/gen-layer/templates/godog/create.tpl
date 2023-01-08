{{define "create"}}
package features

import "github.com/cucumber/godog"

func (s *Suite) {{.CamelCase}}MustBeCreated() error {
       return godog.ErrPending
}

func (s *Suite) userCreate{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}