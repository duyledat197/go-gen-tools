{{define "update"}}
func (u *{{.CamelCase}}Repository) Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	return nil
}
{{end}}