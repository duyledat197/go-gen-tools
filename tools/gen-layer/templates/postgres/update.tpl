{{define "update"}}
func (r *{{.CamelCase}}Repository) Update(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}
{{end}}