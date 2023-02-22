{{define "function"}}
func (s *Suite) {{.FunctionName}} error {
       return godog.ErrPending
}
{{end}}