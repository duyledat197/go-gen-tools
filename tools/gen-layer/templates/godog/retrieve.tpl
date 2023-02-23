{{define "retrieve"}}
{{template "default"}}

func (s *Suite) {{.CamelCase}}IsDeleted() error {
       return godog.ErrPending
}

func (s *Suite) userRetrieve{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}