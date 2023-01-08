{{define "update"}}
package features

import "github.com/cucumber/godog"

func (s *Suite) {{.CamelCase}}IsDeleted() error {
       return godog.ErrPending
}

func (s *Suite) updated{{.PascalCase}}SetAsExpected() error {
       return godog.ErrPending
}

func (s *Suite) userUpdate{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}