{{define "retrieve"}}
func (u *{{.CamelCase}}Repository) GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error) {
	return nil, nil
}
{{end}}