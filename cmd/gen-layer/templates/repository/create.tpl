{{define "create"}}
func (u *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error {
	return nil
}
{{end}}