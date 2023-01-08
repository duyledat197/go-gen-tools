{{define "list"}}
package features

import "github.com/cucumber/godog"

func (s *Suite) returnsStatusCode(arg1 string) error {
       return godog.ErrPending
}

func (s *Suite) userList{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}