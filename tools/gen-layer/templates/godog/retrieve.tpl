{{define "retrieve"}}
package features

import "github.com/cucumber/godog"

func (s *Suite) {{.CamelCase}}IsDeleted() error {
       return godog.ErrPending
}

func (s *Suite) userRetrieve{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}