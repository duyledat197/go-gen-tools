{{define "update"}}
func (u *{{.CamelCase}}Repository) Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {
	return nil
}
{{end}}