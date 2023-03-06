{{define "list"}}
func (r *{{.CamelCase}}Repository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {
	q := models.New(r.db)
	return nil, nil
}
{{end}}