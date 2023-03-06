{{define "retrieve"}}
func (r *{{.CamelCase}}Repository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.{{.PascalCase}}, error) {
	q := models.New(r.db)
	return nil, nil
}
{{end}}