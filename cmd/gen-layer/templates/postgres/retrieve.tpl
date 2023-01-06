{{define "retrieve"}}
func (u *{{.CamelCase}}Repository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.{{.PascalCase}}, error) {
	return nil, nil
}
{{end}}