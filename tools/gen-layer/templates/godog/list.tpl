{{define "list"}}
package features

import "github.com/cucumber/godog"

func returnsStatusCode(arg1 string) error {
       return godog.ErrPending
}

func userList{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}