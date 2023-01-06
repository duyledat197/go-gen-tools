{{define "create"}}
func (u *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	return nil
}
{{end}}