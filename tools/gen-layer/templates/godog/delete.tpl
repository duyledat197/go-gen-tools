{{define "delete"}}
{{template "default"}}

func (s *Suite) {{.CamelCase}}HaveBeenDeletedCorrectly() error {
       return godog.ErrPending
}

func (s *Suite) userDelete{{.PascalCase}}() error {
       return godog.ErrPending
}

func (s *Suite) userDelete{{.PascalCase}}Again() error {
       return godog.ErrPending
}
{{end}}