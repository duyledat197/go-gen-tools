{{define "update"}}
package features

import "github.com/cucumber/godog"

func {{.CamelCase}}IsDeleted() error {
       return godog.ErrPending
}

func updated{{.PascalCase}}SetAsExpected() error {
       return godog.ErrPending
}

func userUpdate{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}