{{define "create"}}
func (r *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}
{{end}}