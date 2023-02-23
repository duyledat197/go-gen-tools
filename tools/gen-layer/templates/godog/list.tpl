{{define "list"}}
{{template "default"}}

func (s *Suite) returnsStatusCode(arg1 string) error {
       return godog.ErrPending
}

func (s *Suite) userList{{.PascalCase}}() error {
       return godog.ErrPending
}
{{end}}