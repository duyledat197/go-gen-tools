{{define "retrieve"}}
package features

import "github.com/cucumber/godog"

func {{.CamelCase}}IsDeleted() error {
       return godog.ErrPending
}

func userRetrieve{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}